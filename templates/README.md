# goenv

![github workflow](https://github.com/drewgonzales360/goenv/actions/workflows/github-actions.yml/badge.svg)

goenv is an small, simple binary that executes the [install instructions](https://go.dev/doc/install) on the Go website and manages several Go versions. There are several other implementations that have much more support, but this has fewer features by design.

## Install

See the "Releases" page and download the latest release.

```bash
# Linux
curl -sSL https://github.com/drewgonzales360/goenv/releases/download/${SEMVER}/goenv-linux-amd64-${SEMVER}.tar.gz -o /tmp/goenv-${SEMVER}.tar.gz

# Mac
curl -sSL https://github.com/drewgonzales360/goenv/releases/download/${SEMVER}/goenv-darwin-amd64-${SEMVER}.tar.gz -o /tmp/goenv-${SEMVER}.tar.gz

tar -xzvf /tmp/goenv-${SEMVER}.tar.gz -C /tmp
mv /tmp/goenv /usr/local/bin
```

## Usage

Calling `goenv` without any arguments will print out a helpful block of text, but here are a few useful examples.

```bash
# Install and use a go version
goenv install 1.14

# Use an installed version. This will fail if you don't have it installed 😥
goenv use 1.17

# Removes an installed version
goenv uninstall 1.16

# Lists installed versions
goenv list
```
