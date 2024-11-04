# Contributing

### Prerequisites

* [Go](https://golang.org/doc/install)
* [Cobra CLI](https://cobra.dev)
* [Docker](https://docs.docker.com/get-docker/)
* [Docker Compose](https://docs.docker.com/compose/install/)

### Run development environment

```bash
go run main.go --help
```

### Add new command

```bash
cobra add <command>
```

### Generate documentation

```bash
cobra gen docs
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
unset GITLAB_TOKEN
goreleaser release --clean
```
