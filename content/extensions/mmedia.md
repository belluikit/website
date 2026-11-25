---
title: Media
description: Video and audio playback with transport controls, subtitles and camera capture, decoded by FFmpeg.
weight: 5
---

`x/mmedia` plays video and audio inside a Bellui window. It opens files,
readers and network streams, shows play, pause, seek and volume controls,
keeps sound and picture in sync, draws subtitles, and shows live video from a
camera. Audio can also play with no widget at all, so a song keeps playing
while the app changes screens. Decoding is done by
[FFmpeg](https://ffmpeg.org), through the
[`goffmpeg`](https://github.com/belluikit/goffmpeg) binding, which brings
prebuilt FFmpeg libraries with it.

## Install

```sh
go get github.com/belluikit/bellui/x/mmedia@v1.5.0
```

There is no FFmpeg to install. `x/mmedia` v1.5.0 requires `goffmpeg`
v0.13.1, which brings the FFmpeg libraries for your platform as a Go module
(`goffmpeg-libs` v0.10.0). They are published for macOS, Linux and Windows on
amd64 and arm64, Android, iOS and the iOS Simulator.

FFmpeg comes as shared libraries, not linked into your executable. A user can
then replace them with their own build, which FFmpeg's licence asks for (see
[Licensing](#licensing)). It also means your app has to carry them:

- **`go run` and `go test`** find the libraries in the Go module cache on
  macOS and Linux. On Windows, put the library module's `shared` folder on
  `PATH` first (`shared-audio` for the audio set, below). `go list -m -f
  '{{.Dir}}' github.com/belluikit/goffmpeg-libs/windows_amd64` prints where
  the module is.
- **A binary from `go build`** looks for the libraries in your module cache,
  so it runs only on the machine that built it.
- **[`bellui package`](/docs/packaging/)** copies the libraries your app loads
  into the package and points the app at them: `Contents/Frameworks` in a
  macOS app, `lib/<app ID>/` in a Linux package, beside the `.exe` on
  Windows, each ABI's `lib/<abi>/` in an Android APK, and `Frameworks/` in an
  iOS app, where they are signed with it. Use it to ship an app.

Other things to know:

- **Linux:** the FFmpeg libraries use the system's OpenSSL 3 for HTTPS
  (`libssl.so.3` and `libcrypto.so.3`), and need glibc 2.35 or newer. `bellui
  package` does not bundle OpenSSL. Sound goes to PulseAudio, or to ALSA when
  no PulseAudio server is running; neither needs a development package.
- If your shell exports `CGO_CFLAGS` or `CGO_LDFLAGS`, for example pointing at
  Homebrew, they can put another FFmpeg ahead of the bundled one. This matters
  most when cross-compiling; clear them if a build picks up the wrong FFmpeg.

### Formats

The FFmpeg libraries come in two codec sets. A build tag picks one:

| Set | Plays | Build |
| --- | --- | --- |
| Media, on desktop | Every format and codec FFmpeg decodes, and cameras | default |
| Media, on Android and iOS | MP3, WAV, MP4 (`.mp4`, `.m4a`, `.mov`), MPEG-TS, HLS, WebM and Matroska, and Ogg, with H.264, VP9 or AV1 video and AAC or Opus audio | default |
| Audio | MP3, WAV, AAC (`.m4a`, `.mp4`, ADTS) and Opus (`.ogg`, `.opus`). No video and no cameras | `-tags goffmpeg_novideo` |

On Android and iOS, H.264 and AAC are decoded only by the platform's own
decoders. A stream they refuse, such as 10-bit H.264, does not play on a
phone. Other codecs, such as FLAC, Vorbis and HEVC, are not in the phone set.
A file outside the set fails to open with
`goffmpeg.ErrUnrecognizedInput`, or with `goffmpeg.ErrNoDecoder` when the
container opens but none of its streams can be decoded.

The audio set is much smaller: on macOS arm64, 2.4 MB of FFmpeg against
15.9 MB for the media set. Built with `goffmpeg_novideo`, a package carries
only the audio set.

FFmpeg's own log is set to errors only, so its warnings do not appear on
your app's stderr. `goffmpeg.SetLogLevel(goffmpeg.LogWarning)` brings them
back while you debug; the level applies to the whole process.

Read [Licensing](#licensing) before you ship an app that uses this module.

## Playing a file or stream

{{< screenshot name="player" alt="The media player showing a test-pattern video with the word Bellui, and its control bar below: pause, stop, elapsed and total time, seek bar, mute and volume." >}}

{{< example file="extensions/mmedia/player/main.go" region="player" >}}

[`MediaPlayer`](https://pkg.go.dev/github.com/belluikit/bellui/x/mmedia#MediaPlayer)
is one widget: the picture, kept at its aspect ratio, with a bar of controls
below it for play and pause, stop, seeking, mute and volume, and the elapsed
and total time. Call `Close` when you are done with it to release the decoder.
Nothing else closes a player: taking it out of a window, or Bellui rebuilding
its renderer, leaves it playing.

There are two ways to open media, and they are not interchangeable:

- **`NewMediaPlayerURL(path or URL)`** hands the location to FFmpeg, which does
  the reading. Use it for files on disk and for anything on the network: it can
  seek within an HTTP stream, and it is the only form that plays formats that
  fetch their own data, such as HLS playlists. HTTP, HTTPS, HLS, RTSP and RTMP
  are supported. On Android, FFmpeg has no TLS, so an `https` file is read
  through Go instead, with the system's certificates; it still streams and
  seeks, but an `https` HLS playlist does not open there.
- **`NewMediaPlayer(io.Reader)`** reads through Go, for media you already have
  as a reader, such as an embedded file or a decrypted stream. The format is
  detected automatically. If the reader also implements `io.Seeker`, the user
  can seek; otherwise playback only goes forward.

`WithScrubPreview` shows a thumbnail of the frame under the pointer while it
hovers over the seek bar. It opens the media a second time, so it works with
the URL form, or with an `*os.File` passed to the reader form, and the
constructor returns an error if you ask for it with any other reader.

### Options

Every constructor takes the same options:

| Option | Effect |
| --- | --- |
| `WithAutoPlay()` | Start playing as soon as the media is open. |
| `WithLoop()` | Start again from the beginning at the end. |
| `WithVolume(v)` | Initial volume, from 0.0 to 1.0. |
| `WithScrubPreview()` | Preview frames on the seek bar, as above. |
| `WithOnTimeUpdate(fn)` | Called with the current position and the duration: every 100 ms while playing, and once on each pause, each seek and at the end. Nothing arrives while paused. |
| `WithOnSeek(fn)` | Called once for every seek, with the position before and after it. |
| `WithOnEnd(fn)` | Called when playback reaches the end. |
| `WithOnError(fn)` | Called when playback fails, including a seek that fails. |
| `WithDownmix()` | Deprecated, and has no effect. Surround audio is mixed down to stereo without it; see [Audio output](#audio-output). |

`WithOnSeek` reports a seek from the seek bar (once, when the user lets go)
and one your code asks for with `Seek`. It runs before the time update for
the new position, so you can tell that update from normal playback.
The rewind that `Stop` does is not a seek, and neither is a loop starting
again. A live source cannot seek.

Since time updates stop while nothing plays, an app that used them as a
regular tick needs a timer of its own.

{{< callout type="warning" >}}
`WithOnTimeUpdate`, `WithOnEnd` and `WithOnError` run on the player's own
goroutine, not the UI thread. `WithOnSeek` runs on the goroutine that called
`Seek`, which is the UI thread for the seek bar. Wrap anything that touches
widgets in `bellui.Do`, as the examples do.
{{< /callout >}}

From code, `Play`, `Pause`, `Stop`, `Seek`, `SetVolume` and `SetLoop` control
playback, and `CurrentTime`, `Duration`, `IsPlaying` and `VideoSize` report on
it. `Stats` returns frame counts and how far the picture is from the sound,
which is useful when diagnosing stutter.

## Audio only, video only

[`AudioPlayer`](https://pkg.go.dev/github.com/belluikit/bellui/x/mmedia#AudioPlayer)
is the same player without a picture: just the controls. `VideoPlayer` is the
reverse: a picture without controls, for when you build your own.

{{< example file="extensions/mmedia/audio/main.go" region="audio" >}}

`NewTransportControls` gives you the control bar on its own, to drive a player
of your own design through its `OnPlay`, `OnPause`, `OnSeek` and `OnVolume`
fields.

### Playback that outlives a screen

An `AudioPlayer` built with `NewAudioPlayer` owns its sound: closing the
widget stops it. For music that keeps playing while the user moves between
screens, open an
[`AudioPlayback`](https://pkg.go.dev/github.com/belluikit/bellui/x/mmedia#AudioPlayback)
instead. It plays, pauses, seeks, stops and loops with no widget, window or
app, and takes the same options as the players. Give each screen a view of it
with `NewAudioPlayerFor`:

{{< example file="extensions/mmedia/playback/main.go" region="playback" >}}

Attaching a view neither starts, stops nor rewinds the playback. Closing the
view, or calling its `Detach`, leaves the playback as it was. Your app closes
the `AudioPlayback` when it is done with it. `AudioPlayer.Playback` returns
the playback a widget shows. `AudioPlayback.Level` reports how loud the sound
is right now, from 0 to 1, cheap enough to call once a frame for a
visualizer.

The `WithOnTimeUpdate`, `WithOnEnd` and `WithOnError` callbacks of an
`AudioPlayback` must not wait for its `Stop` or `Close`: that deadlocks.

## Subtitles

A [`SubtitleOverlay`](https://pkg.go.dev/github.com/belluikit/bellui/x/mmedia#SubtitleOverlay)
shows timed text over the picture. It reads SubRip (`LoadSRT`) and WebVTT
(`LoadWebVTT`) files. `MediaPlayerWithSubtitles` lays it over a player and
keeps it in step with playback:

{{< example file="extensions/mmedia/subtitles/main.go" region="subtitles" >}}

`MediaPlayerWithSubtitles` takes over the player's time updates, so a
`WithOnTimeUpdate` callback given to that player stops being called. If you
need both, place the overlay yourself and call its `SetPosition` from your own
time callback.

To show subtitles from another source, such as a stream's own subtitles or
a format the package does not parse, build a
[`SubtitleTrack`](https://pkg.go.dev/github.com/belluikit/bellui/x/mmedia#SubtitleTrack)
of `SubtitleCue`s and pass it to `LoadTrack`:

{{< example file="extensions/mmedia/cues/main.go" region="cues" >}}

A cue shows from `Start` to `End`, both included. Do not change a track after
you load it: the overlay reads it on every position update.

## Cameras

[`NewCapturePlayer`](https://pkg.go.dev/github.com/belluikit/bellui/x/mmedia#NewCapturePlayer)
opens a camera instead of a recording and returns the same `MediaPlayer`
widget. The player is live: it has no duration, the seek bar shows "live" and
is disabled, and `WithLoop` has no effect. `Stop` releases the camera, so its
light goes off and other apps can use it; `Play` takes it back.

{{< example file="extensions/mmedia/capture/main.go" region="capture" >}}

[`CaptureSelect`](https://pkg.go.dev/github.com/belluikit/bellui/x/mmedia#CaptureSelect)
is a drop-down list of cameras. Its callback runs only when the user picks one.
`SetCaptureDevice` switches the running player to another camera without
rebuilding the widget. It opens the new camera before releasing the old one, so
if the switch fails the old picture stays. Opening a camera takes a noticeable
time, so call it from a goroutine, as above.

How cameras are found differs by platform:

- **Linux** (V4L2) and **Windows** (DirectShow) can list their cameras, and
  `VideoCaptureDevices` returns them. Use `Label()` for the name to show the
  user, and `Name` to open the device.
- **macOS** (AVFoundation) cannot list cameras. They are chosen by position
  instead: `NewCapturePlayerIndex(0)` is the first. On macOS `CaptureSelect`
  starts empty and shows a button that finds cameras by trying to open each
  position in turn, which takes about a second per position and briefly turns
  on the camera light. Some of the positions it finds are screen-capture
  inputs, not cameras.

`WithCaptureOptions("framerate=30,pixel_format=nv12")` passes settings to the
platform's capture driver; the accepted keys differ per platform.

There is no camera capture on Android or iOS, or in the audio set. A camera
produces frames only as they happen, so a capture player keeps only three
frames ready ahead of the picture, which keeps the delay and the memory use
low.

## Audio output

A process has one audio device, opened by the first player with audio and
never reopened. It plays 48 kHz stereo, whatever the stream. Each stream is
converted to it: 5.1 and 7.1 are mixed down to stereo, and mono plays on both
channels, each 3 dB quieter than the source.

To get surround output, call `SetAudioDeviceChannels` before the first
player with audio is built:

{{< example file="extensions/mmedia/surround/main.go" region="surround" >}}

After the device has opened, the call returns an error unless the count is
the one in use. If the output refuses the count, every player with audio
fails with the device's error until the app restarts, so ask for surround
only where the output is known to take it.

A player with audio fails in its constructor when the device will not open,
for example on a machine with no sound output, or when the stream cannot be
converted to the device's format. After one such failure, every later player
with audio returns the same error. A video with no audio stream never opens
the device.

### In the background on iOS

On iOS, `x/mmedia` sets the app's audio session to the playback category.
Audio then plays through the silent switch, and keeps playing when the app
is in the background or the screen locks, as a music app does. iOS allows
the background part only when the app declares it, in `ios/Info.plist` next
to `manifest.pxf`:

```xml
<key>UIBackgroundModes</key>
<array>
	<string>audio</string>
</array>
```

## Renderers

Video plays under both [renderers](/develop/rendering/), the default WebGPU
one and the OpenGL one (`-tags gl`). Under both, each new frame updates the
picture's GPU texture in place rather than creating a new one. The decoder
converts each frame straight into a set of frames allocated once, so
playback allocates no memory per frame.

## Licensing

`x/mmedia` itself is BSD 3-Clause licensed, like Bellui, and so is the
`goffmpeg` binding's own code. **FFmpeg is not.** The FFmpeg libraries that
`goffmpeg` links are licensed under the
[GNU Lesser General Public License, version 2.1 or later](https://www.gnu.org/licenses/old-licenses/lgpl-2.1.html).
They are built without FFmpeg's GPL and non-free components, so no stricter
licence applies.

When you build an app with `x/mmedia`, it loads the FFmpeg libraries as
shared libraries, and `bellui package` ships them inside your app. Your app
is then a "work that uses the Library" in the LGPL's terms, and **when you
distribute it, the LGPL 2.1 applies.** In outline, it asks you to:

- say prominently that your app uses FFmpeg and that FFmpeg is covered by the
  LGPL, and include a copy of the licence;
- make FFmpeg's source code available for the version you ship;
- let users replace FFmpeg with a modified version. Shipping FFmpeg as
  separate shared libraries, as `bellui package` does, is the way section 6
  of the licence provides for this;
- allow modification of your app for the user's own use, and reverse
  engineering to debug such modifications.

This is a summary, not legal advice. Read the licence, sections 4 and 6
above all, and ask a lawyer if you are unsure what it requires of you.

The details you need for the notice are published with each platform's
library module in the
[`goffmpeg-libs`](https://github.com/belluikit/goffmpeg-libs) repository: the
FFmpeg version, the exact git revision, the configuration of both codec sets,
and a link to the matching FFmpeg source. The `goffmpeg`
[`NOTICE`](https://github.com/belluikit/goffmpeg/blob/main/NOTICE) file covers
the FFmpeg headers it contains.
