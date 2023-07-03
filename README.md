# Thinkport CLI

This cli tool provides a set of commands related to Thinkport.
This is a technical experiement and not intended for production use.

## Features

* [x] List all Thinkport members or a single member

### Planned

* [ ] List all Thinkport projects
* [ ] Display cooperate identity information
* [ ] Display contact information

## Installation

```bash
brew install thinkport/tap/thinkport
```

## Usage

```bash
thinkport --help
```

## Development

### Prerequisites

* [Go](https://golang.org/doc/install)
* [Docker](https://docs.docker.com/get-docker/)
* [Docker Compose](https://docs.docker.com/compose/install/)

### Run development environment

```bash
encore run
```

### Test local build

```bash
goreleaser build --snapshot --clean
chmod -R u+x dist
 ./dist/thinkport_darwin_arm64/thinkport --help
```

### Test

```bash
encore test ./...
```

### Release

Releases are made with goreleaser. To create a new release, run:

```bash
goreleaser release --clean
```

```bash
