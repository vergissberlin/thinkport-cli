# 🌈 Thinkport CLI

This cli tool provides a set of commands related to Thinkport.
This is a technical experiement and not intended for production use.

## 🚀 Features

* [x] List all Thinkport members or a single member
* [ ] List all Thinkport products or a single product
* [ ] Display cooperate identity information
* [ ] Display contact information
* [ ] Automatic update check
* [ ] Automatic update installation
* [ ] Automatic update notification
* [ ] Count members
* [ ] Display news

## 📀 Installation

### Homebrew (macOS and Linux)

```bash
brew install vergissberlin/tap/thinkport
```

### Manual (Windows, macOS, Linux)

Download the latest release from the [releases page](https://github.com/vergissberlin/thinkport/releases) and install it manually.

## 👩‍💻 Usage

```bash
thinkport --help
thinkport members
```

This will list all Thinkport members.

## 👷‍♀️ Development

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

```bash
