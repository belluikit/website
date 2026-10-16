---
title: Advanced topics
description: Sign-in, analytics, web content, developer tools, native views, animation and drag and drop.
weight: 14
---

These pages cover features most apps need only some of the time. Several of
them are separate packages. A package is compiled into your binary only when
you import it, so an app that never embeds web content or sends analytics
carries none of that code.

{{< cards >}}
  {{< card link="auth/" title="Sign-in with OIDC and OAuth" subtitle="Branded sign-in buttons and the system-browser flow." >}}
  {{< card link="analytics/" title="Analytics" subtitle="Opt-in, consent-first event tracking with batching and offline queueing." >}}
  {{< card link="webview/" title="Web views" subtitle="Embed web content and talk to it from Go." >}}
  {{< card link="devtools/" title="Developer tools" subtitle="The layout overlay, tree dump and the widget inspector." >}}
  {{< card link="native-views/" title="Native views" subtitle="Place an OS-owned view inside a Bellui layout." >}}
  {{< card link="animation/" title="Animation" subtitle="Frame callbacks, easing curves, springs and keyframes." >}}
  {{< card link="drag-and-drop/" title="Drag and drop" subtitle="Inside the app, out to other apps, and files in from the OS." >}}
{{< /cards >}}
