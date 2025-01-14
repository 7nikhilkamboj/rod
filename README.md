# Overview

[![Go Reference](https://pkg.go.dev/badge/github.com/7nikhilkamboj/rod.svg)](https://pkg.go.dev/github.com/7nikhilkamboj/rod)
[![Discord Chat](https://img.shields.io/discord/719933559456006165.svg)][discord room]

## [Documentation](https://go-rod.github.io/) | [API reference](https://pkg.go.dev/github.com/7nikhilkamboj/rod?tab=doc) | [Management](https://github.com/orgs/go-rod/projects/1)

Rod is a high-level driver directly based on [DevTools Protocol](https://chromedevtools.github.io/devtools-protocol).
It's designed for web automation and scraping for both high-level and low-level use, senior developers can use the low-level packages and functions to easily
customize or build up their own version of Rod, the high-level functions are just examples to build a default version of Rod.

## Features

- Chained context design, intuitive to timeout or cancel the long-running task
- Auto-wait elements to be ready
- Debugging friendly, auto input tracing, remote monitoring headless browser
- Thread-safe for all operations
- Automatically find or download [browser](lib/launcher)
- Lightweight, no third-party dependencies, [CI](https://github.com/7nikhilkamboj/rod/actions) tested on Linux, Mac, and Windows
- High-level helpers like WaitStable, WaitRequestIdle, HijackRequests, WaitDownload, etc
- Two-step WaitEvent design, never miss an event ([how it works](https://github.com/ysmood/goob))
- Correctly handles nested iframes or shadow DOMs
- No zombie browser process after the crash ([how it works](https://github.com/ysmood/leakless))

## Examples

Please check the [examples_test.go](examples_test.go) file first, then check the [examples](lib/examples) folder.

For more detailed examples, please search the unit tests.
Such as the usage of method `HandleAuth`, you can search all the `*_test.go` files that contain `HandleAuth`,
for example, use Github online [search in repository](https://github.com/7nikhilkamboj/rod/search?q=HandleAuth&unscoped_q=HandleAuth).
You can also search the GitHub [issues](https://github.com/7nikhilkamboj/rod/issues) or [discussions](https://github.com/7nikhilkamboj/rod/discussions),
a lot of usage examples are recorded there.

[Here](lib/examples/compare-chromedp) is a comparison of the examples between rod and Chromedp.

If you have questions, please raise an [issues](https://github.com/7nikhilkamboj/rod/issues)/[discussions](https://github.com/7nikhilkamboj/rod/discussions) or join the [chat room][discord room].

## Join us

Your help is more than welcome! Even just open an issue to ask a question may greatly help others.

Please read [How To Ask Questions The Smart Way](http://www.catb.org/~esr/faqs/smart-questions.html) before you ask questions.

We use Github Projects to manage tasks, you can see the priority and progress of the issues [here](https://github.com/orgs/go-rod/projects/1).

If you want to contribute code for this project please read the [Contributor Guide](.github/CONTRIBUTING.md).

[discord room]: https://discord.gg/CpevuvY
[中文文档]: https://go-rod.github.io/i18n/zh-CN
