# Contributing


### Prerequisites

* [Go](https://golang.org/doc/install)
* [Docker](https://docs.docker.com/get-docker/)
* [Docker Compose](https://docs.docker.com/compose/install/)

### Run development environment

```bash
encore run
```

### 🧪 Test local build

```bash
goreleaser build --snapshot --clean
chmod -R u+x dist
 ./dist/thinkport_darwin_arm64/thinkport --help
```

### 🧪 Test

```bash
encore test ./...
```

### Release

Releases are made with goreleaser. To create a new release, run:

```bash
goreleaser release --clean
```
