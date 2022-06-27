# goenv

![github workflow](https://github.com/drewgonzales360/goenv/actions/workflows/github-actions.yml/badge.svg)

goenv is an small, simple binary that executes the [install instructions](https://go.dev/doc/install) on the Go website and manages several Go versions. There are several other implementations that have more features, but this has fewer ones by design. Goenv downloads and extracts go to `/usr/local/goenv/<VERSION>` and adds a symlink from `/usr/local/go -> /usr/local/goenv/<VERSION>`. It doesn't do anything fancy like guess the version of a project, intercept calls to `go`, or check for new versions.

## Usage

Calling `goenv` without any arguments will print out a helpful block of text, but here are a few useful examples.

```bash
# Install and use a go version
goenv install 1.14

# Use an installed version. This will fail if you don't have it installed.
goenv use 1.17

# Removes an installed version
goenv uninstall 1.16

# Lists installed versions
goenv list
```

## Install

To install goenv, follow the steps below. Older releases are in the Releases page.

```bash
# Step 1: Linux Only
curl -sSL https://github.com/drewgonzales360/goenv/releases/download/v0.0.3/goenv-linux-amd64-v0.0.3.tar.gz -o /tmp/goenv-v0.0.3.tar.gz

# Step 1: Mac Only
curl -sSL https://github.com/drewgonzales360/goenv/releases/download/v0.0.3/goenv-darwin-amd64-v0.0.3.tar.gz -o /tmp/goenv-v0.0.3.tar.gz

# Step 2: Extract and Install Go
tar -xzvf /tmp/goenv-v0.0.3.tar.gz -C /tmp
mv /tmp/goenv /usr/local/bin

# Step 3: Add /usr/local/go/bin to PATH or put this line in whichever dotfile is used
export PATH=/usr/local/go/bin:PATH
```

It's best to install this binary without `go install` so that it is managed independent of Go.

## Configuration

```json
{
    "go.goroot": "/Users/drew.gonzales/.local/go"
}
```

If your VScode editor throws you weird errors on start up.
