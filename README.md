# goenv

![github workflow](https://github.com/drewgonzales360/goenv/actions/workflows/github-actions.yml/badge.svg)

goenv is an small, simple binary that executes the [install instructions](https://go.dev/doc/install) on the Go website and manages several Go versions. Goenv downloads and extracts go to `/usr/local/goenv/<VERSION>` and adds a symlink from `/usr/local/go -> /usr/local/goenv/<VERSION>`. It was heavily inspired by [Dave Cheney's blog post](https://dave.cheney.net/2014/04/20/how-to-install-multiple-versions-of-go).

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

| Environment Variable  | Default             | Explanation |
| -                     | -                   | - |
| GOENV_ROOT_DIR        | "/usr/local/go"     | This is the default Go root. This is should be in your path and links to the GOENV_INSTALL_DIR |
| GOENV_INSTALL_DIR     | "/usr/local/goenv"  | Directory where your various Go installations will be installed |

The default `GOROOT` usually requires root access. You can avoid it by setting the configuration variables above. Adding

```shell
# goenv configuration
export GOENV_INSTALL_DIR="/home/drew/.local/goenv"
export GOENV_ROOT_DIR="/home/drew/.local/go"
export PATH="/mnt/NRG/usr/local/go/bin:/mnt/NRG/usr/local/goenv/bin:/mnt/NRG/usr/local/go/bin:/mnt/NRG/usr/local/go/bin:/home/drew/.cargo/bin:/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin:/usr/games:/usr/local/games:/snap/bin:/snap/bin:/home/drew/.fzf/bin:/mnt/NRG/Docs/CS/Code/go/bin:/home/drew/.cargo/bin:/mnt/NRG/Docs/CS/Code/go/bin:/home/drew/.cargo/bin"
```

If your VScode editor throws you weird errors on start up, add this to your vscode settings.json.
```json
{
    "go.gopath": "/your/gopath",
    "go.goroot": "/your/goroot", # or whatever /mnt/NRG/usr/local/go is set to
}
```
