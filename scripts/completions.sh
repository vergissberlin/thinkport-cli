#!/bin/sh
set -e

# Check if go is installed
if ! command -v go >/dev/null; then
    echo "go is not installed. Please install it with:"
    echo "https://golang.org/doc/install"
    exit 1
fi

# Check if cobra-cli is installed
if ! command -v cobra-cli >/dev/null; then
    echo "cobra-cli is not installed. Please install it with:"
    echo "go get -u github.com/spf13/cobra/cobra"
    exit 1
fi

PATH_COMPLETIONS=completions
rm -rf $PATH_COMPLETIONS
mkdir -p $PATH_COMPLETIONS
for sh in bash zsh fish; do
	cobra-cli completion "$sh" >"$PATH_COMPLETIONS/thinkort.$sh"
done
