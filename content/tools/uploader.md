---
title: uploader
description: A resumable file transfer service over gRPC, with a Go client SDK. Transfers that are cut off continue from where they stopped.
weight: 1
---

uploader moves large files between a client and a server. The client sends
the file in chunks. If the connection drops, the client asks the server how
much it already has and sends only the rest. Downloads work the same way in
the other direction.

Use it when your app sends or receives files that are too big to retry from
the start: videos, data sets, backups. It is a separate service. Your Bellui
app talks to it through the Go SDK, the same way any Go program would.

- Repository: [github.com/belluikit/uploader](https://github.com/belluikit/uploader)
- License: BSD-3-Clause
- Current release: [v0.1.0](https://github.com/belluikit/uploader/releases/tag/v0.1.0)

## What is in it

| Part | What it is |
| --- | --- |
| `cmd/uploader` | The server binary |
| `sdk` | The Go client |
| `proto/uploader/v1` | The `TransferService` gRPC definition and its generated Go code |
| `server` | The service implementation, if you want to build your own server binary |
| `storage` | The storage interface, and a local-disk implementation |
| `session` | Transfer sessions, saved to disk |

It is pure Go with no build-tagged files. You need Go 1.26.1 or later.

## Install

The server:

```sh
go install github.com/belluikit/uploader/cmd/uploader@latest
```

The SDK, in your app's module:

```sh
go get github.com/belluikit/uploader/sdk@latest
```

Get the `sdk` package, not the bare module path. The module has no package at
its root, so `go get github.com/belluikit/uploader` leaves out the gRPC
dependencies the SDK needs, and the build then fails on missing `go.sum`
entries.

## Quick start

{{% steps %}}

### Start a server

```sh
uploader serve --listen :9090 --data-dir data
```

`--listen` is the gRPC address (default `:9090`). `--data-dir` is where files
and sessions go (default `data`). Press Ctrl-C to stop it.

### Send and fetch a file from Go

The repository README has a complete program:
[Using the Go SDK](https://github.com/belluikit/uploader#using-the-go-sdk).
It connects with `sdk.New("localhost:9090")`, uploads a file with
`Client.Upload`, and downloads it again by its file ID with
`Client.Download`.

{{% /steps %}}

{{< callout type="warning" >}}
The `uploader` server listens on plain gRPC. It has no TLS and no
authentication, so anyone who can reach the port can upload and download
files. Run it only where you trust everyone who can reach it.
{{< /callout >}}

## The SDK

[`sdk.New(addr, opts...)`](https://pkg.go.dev/github.com/belluikit/uploader/sdk#New)
connects to a server. With no `grpc.DialOption`s it uses plaintext; pass your
own options to change that. `sdk.NewFromConn` wraps a `*grpc.ClientConn` you
already have.

| Call | Returns |
| --- | --- |
| `Client.Upload(ctx, path, opts...)` | `*UploadResult` and an error. The result has `SessionID`, `FileID`, `CommittedSize` and `Complete`. |
| `Client.Download(ctx, fileID, destPath, opts...)` | `*DownloadResult` and an error. The result has `SessionID`, `Size` and `Complete`. |

Both return their result even when the transfer fails part way. Keep its
`SessionID`: it is what you resume with.

| Option | Effect |
| --- | --- |
| `WithChunkSize(n)` | Upload chunk size in bytes. The default is 256 KB. It must be between 1 and `sdk.MaxChunkSize`, just under the server's 4 MiB message limit, or `Upload` returns `ErrInvalidChunkSize`. |
| `WithProgress(fn)` | Called after each chunk with the bytes sent so far and the total. |
| `WithSessionID(id)` | Resume an earlier upload or download. |
| `WithMaxRetries(n)` | How many times a transfer resumes itself after a temporary failure. The default is 3; 0 turns it off. |

`WithProgress` is called from the goroutine doing the transfer. In a Bellui
app, run the transfer in a goroutine and hand UI updates back with
[`bellui.Do`](https://pkg.go.dev/github.com/belluikit/bellui#Do), or update a
data binding. The gallery's File Transfer page does both.

## How resuming works

**Uploads.** Pass the earlier `SessionID` with `WithSessionID`. The SDK asks
the server about the session first. If the session is unknown or has
expired, or if the local file's SHA-256 hash changed, it starts a new upload.
Otherwise it skips to the point the server has saved and sends the rest.

**Downloads.** Pass the download's session ID and the same `destPath`. The
SDK continues from the current size of the file at `destPath`. If the file on
the server changed since, the server refuses with `FailedPrecondition`.

**Automatic retries.** When a transfer fails with a temporary gRPC error
(`Unavailable`, `ResourceExhausted`, `Aborted`, or `DeadlineExceeded` from the
server), the SDK waits and resumes on its own. The wait starts at 200 ms and
doubles up to 5 s. Cancelling your context stops it.

## The server

Under `--data-dir` the server keeps:

- `staging/`: uploads in progress;
- `files/`: finished uploads, named by file ID;
- `upload/` and `download/`: one session file per transfer.

The original file name is kept in the upload session, not in the stored
file's name. A session expires 24 hours after its last update, and the
server removes expired sessions once an hour. It also registers gRPC server
reflection.

The binary always stores files on local disk. To store them elsewhere,
implement the `storage.Backend` interface and pass it to `server.New` in a
server binary of your own.

## The gRPC API

Any gRPC client can use the service; the Go SDK is one of them.
`TransferService`, in
[`proto/uploader/v1/uploader.proto`](https://github.com/belluikit/uploader/blob/main/proto/uploader/v1/uploader.proto),
has five calls: `StartUpload`, `Upload` (a client stream), `Download` (a
server stream), `QueryUpload` and `QueryDownload`. The README's
[gRPC API](https://github.com/belluikit/uploader#grpc-api) section describes
each one.

## With Bellui

The [gallery app](/gallery/#file-transfer) has a File Transfer page built on
this SDK. You type a server address and connect, pick a file with the file
open dialog to upload it, and save it again with the file save dialog. A
progress bar follows each transfer.
