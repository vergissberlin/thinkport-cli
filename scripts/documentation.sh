#!/bin/sh
set -e

# https://cobra.dev/#generating-documentation-for-your-command

PATH_MAN=docs/manpages

rm -rf $PATH_MAN
mkdir -p $PATH_MAN
go run . man | gzip -c -9 > $PATH_MAN/thinkport.1.gz
