# goenv

![github workflow](https://github.com/drewgonzales360/goenv/actions/workflows/github-actions.yml/badge.svg)

goenv is an small, simple binary that executes the [install instructions](https://go.dev/doc/install) on the Go website and manages several Go versions. Goenv downloads and extracts go to `/usr/local/goenv/<VERSION>` and adds a symlink from `/usr/local/go -> /usr/local/goenv/<VERSION>`. It was heavily inspired by [Dave Cheney's blog post](https://dave.cheney.net/2014/04/20/how-to-install-multiple-versions-of-go).

## Install

To install goenv, follow the steps below. Older releases are in the [Releases page](https://github.com/drewgonzales360/goenv/releases).

```bash
# Step 1: Linux Only
curl -sSL https://github.com/drewgonzales360/goenv/releases/download/v0.0.4/goenv-linux-amd64-v0.0.4.tar.gz -o /tmp/goenv-v0.0.4.tar.gz

# Step 1: Mac Only
curl -sSL https://github.com/drewgonzales360/goenv/releases/download/v0.0.4/goenv-darwin-amd64-v0.0.4.tar.gz -o /tmp/goenv-v0.0.4.tar.gz

# Step 2: Extract and Install Go
tar -xzvf /tmp/goenv-v0.0.4.tar.gz -C /tmp
mv /tmp/goenv /usr/local/bin

# Step 3: Add /usr/local/go/bin (or $GOENV_ROOT_DIR/bin) to PATH
export PATH=/usr/local/go/bin:PATH
```

Install this binary _without_ `go install` so that it is managed independent of Go.

## Usage

Calling `goenv` without any arguments will print out a helpful block of text, but here are a few useful examples. Note that installing 1.14 will install 1.14, even if 1.14.5 is the latest patch version.

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

## Configuration

| Environment Variable  | Default             | Explanation |
| -                     | -                   | - |
| GOENV_ROOT_DIR        | "/usr/local/go"     | This is the default Go root. This is should be in your path and links to the GOENV_INSTALL_DIR |
| GOENV_INSTALL_DIR     | "/usr/local/goenv"  | Directory where your various Go installations will be installed |

The default `GOROOT` usually requires root access. You can avoid it by setting the configuration variables above. Adding

```shell
# goenv configuration
export GOENV_INSTALL_DIR="$HOME/.local/goenv"
export GOENV_ROOT_DIR="$HOME/.local/go"
export PATH="$GOENV_ROOT_DIR/bin:$PATH"
```

If your VScode editor throws you weird errors on start up, add this to your vscode settings.json.
```json
{
    "go.gopath": "/your/gopath",
    "go.goroot": "/your/goroot", # or whatever $GOENV_ROOT_DIR is set to
}
```

## How This Works

`goenv` does the following by default:
  - Downloads the tarball for the corresponding version a user provides. Makes a best effort attempt to check the shasums
  - Extracts the tarball to `/usr/local/goenv/${VERSION}`. For example, `goenv install 1.17.6` will create `/usr/local/goenv/1.17.6`
  - Creates a symlink from `/usr/local/go/` to `/usr/local/goenv/1.17.6/bin/go`. The same thing happens for `gofmt`.

The install directory `/usr/local/goenv` and root directory `/usr/local/go` is configurable through environment variables.

## Other Implementations

There are a few other implementations of this that have more features, but I only needed the binaries. Other implementations can do fancier things like check for the correct Go version for a module and use the corresponding version in runtime, but they require intercepting the call to Go and passing the command to the right version. I didn't want my code to be in the "hot path."

A few other implementations are also written in other languages. This one is written in Go ????.

The recommendation on the Go website is to use your first install of Go to install _other_ versions of Go. Then you'd call other versions of Go like `go1.17.8 build`. This provides a consistent experience.
