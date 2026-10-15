---
title: Files and storage
description: Read and write files through URIs, keep user preferences, and store documents in the app's private area.
weight: 10
---

Bellui gives you three ways to keep data:

- **Files and URIs.** The [`storage`](https://pkg.go.dev/github.com/belluikit/bellui/storage)
  package reads, writes and lists resources named by a URI, such as
  `file:///home/ada/notes.txt`. The same functions work with files picked in a
  [file dialog](/docs/dialogs/files/), with Android content URIs, and with
  storage you plug in yourself.
- **Preferences.** Small typed values (numbers, strings, flags, lists) that are
  saved for you between runs.
- **App storage.** A private folder of documents that belongs to your app.

Preferences and app storage are kept per app, so your app needs a unique ID.
Create the app with
[`app.NewWithID`](https://pkg.go.dev/github.com/belluikit/bellui/app#NewWithID).
(`app.New` can take the ID from your app's [manifest](/docs/packaging/metadata/),
but in this release only when the manifest is next to the program, as with
`go run`.)

{{< cards >}}
  {{< card link="files/" title="Files and URIs" subtitle="Create URIs, read, write, list, copy and delete." >}}
  {{< card link="preferences/" title="Preferences" subtitle="The generic Pref, PrefOr and SetPref functions." >}}
  {{< card link="app-storage/" title="App storage" subtitle="Documents in the app's own folder, and where it lives on each platform." >}}
  {{< card link="repositories/" title="Custom repositories" subtitle="Teach the storage functions a new URI scheme." >}}
{{< /cards >}}
