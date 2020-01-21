# jk

[![Build Status](https://travis-ci.org/jkcfg/jk.svg?branch=master)](https://travis-ci.org/jkcfg/jk)

Caveat: jk is pre-1.0, and as such, features, commands, and APIs inevitably
will change.

Releases within a minor version should be stable; we try to signal breaking
changes ahead of time in [./docs/deprecations.md](./docs/deprecations.md).

## jk - configuration as code

`jk` is a data templating tool designed to help writing structured
configuration files.

The main idea behind `jk` is to use a general purpose language for this task.
They offer mature tooling, great runtimes, a well established ecosystem and
many learning resources. `jk` uses Javascript and a runtime tailored for
configuration.

## Quick start

A good way to start with `jk` is to read our [introduction tutorial][quick-start].
For more context head to our [introduction blog post][blog-0]!

## More complex examples

- [A Kubernetes deployment written in Typescript][guestbook-ts]
- [Generating Kubernetes objects, using e.g., kustomize-like behavior][k8s-examples]
- [Using jk with JavaScript libraries][js-examples]

## Architecture & design

### v8

`jk` itself is a Javascript runtime written in Go and embedding [v8][v8]. It
uses Ryan Dahl's [v8worker2][v8worker2] to embed v8 and
[flatbuffers][flatbuffers] for the v8 ⟷ Go communication.

### Hermeticity

While a general purpose language is great, configuration code can be made
more maintainable by restricting what it can do. A nice property `jk` has to
offer is being *hermetic*: if you clone a git repository
and execute a `jk` script, the resulting files should be the same on any
machine. To give concrete examples, this means the `jk` standard library
doesn't support environment variables nor has any networking capability.

### Library support

`jk` provides an unopinionated data templating layer. On top of the `jk`
runtime, libraries provide APIs for users to write configuration.

- [kubernetes][kubernetes]: build Kubernetes objects
- [grafana][grafana]: build dashboard configs for Grafana

## Roadmap

This project is still in early stages but future (exciting!) plans include:

- Reach the state of having Kubernetes examples working and well documented.
- Work on hermeticity. (eg. [#110][issue110], [#44][issue44], [topic/hermeticity][issueHermeticity]).
- Native typescript support ([#54][issue54]).
- HCL support ([#94][issue94]).

[v8]: https://v8.dev/
[blog-0]: https://damien.lespiau.name/posts/2019-06-12-jk-configuration-as-code/
[quick-start]: https://jkcfg.github.io/#/documentation/quick-start
[grafana]: https://github.com/jkcfg/grafana
[kubernetes]: https://github.com/jkcfg/kubernetes
[guestbook-ts]: https://github.com/jkcfg/kubernetes/blob/master/examples/guestbook-ts
[js-examples]: https://github.com/jkcfg/jk/tree/master/examples
[k8s-examples]: https://github.com/jkcfg/jk/tree/master/examples/kubernetes
[v8worker2]: https://github.com/ry/v8worker2
[flatbuffers]: https://github.com/google/flatbuffers

[issue44]: https://github.com/jkcfg/jk/issues/44
[issue54]: https://github.com/jkcfg/jk/issues/54
[issue94]: https://github.com/jkcfg/jk/issues/94
[issue110]: https://github.com/jkcfg/jk/issues/110
[issueHermeticity]: https://github.com/jkcfg/jk/issues?q=is%3Aissue+is%3Aopen+label%3Atopic%2Fhermeticity

## Contributing
When contributing to this repository, please first discuss the change you wish to make via issue with the owners of this repository before making a change.

### Setup
The `jk` executable itself is written in [Go](https://golang.org), but the JavaScript part of this project requires [NodeJS](https://nodejs.org).

**Prerequisites:**
* `go` 1.11.4 or later (modules support)
* `nodejs`, `npm`
* `make`
* `pkg-config`

First off, clone this repository:
```bash
$ git clone https://github.com/jkcfg/jk.git
# or with hub:
$ hub clone jkcfg/jk

$ cd ./jk
```

Then pull most of the dependencies using the `Makefile`:
```bash
$ make dep
```

`jk` is linked against [`v8worker2`](https://github.com/jkcfg/v8worker2). As building V8 takes ~30 minutes, [prebuilt versions](https://github.com/jkcfg/prebuilt) are provided for `linux/amd64` and `darwin/amd64`. This also includes `flatc` from `flatbuffers`:
```bash
# download the prebuilt artifacts from GitHub:
$ git clone https://github.com/jkcfg/prebuilt.git
$ cd ./prebuilt

# these following put files in /usr/local, so you may need sudo

# x64 Linux:
$ make install-linux-amd64

# x64 macOS:
$ make install-darwin-amd64
```

### Building
After setting up the environment, the `jk` binary can be built:
```bash
$ make jk
$ ./jk --help # confirm it works
```

Additionally, on Linux, it's possible to use a docker container to build the project instead of installing the prebuilt libraries and binaries:
```bash
$ ./run-in-docker.sh make dep jk
```
