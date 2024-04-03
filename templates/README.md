# goenv

![github workflow](https://github.com/drewgonzales360/goenv/actions/workflows/github-actions.yml/badge.svg)
[![Go Reference](https://pkg.go.dev/badge/pkg.go.dev/drewgonzales360/goenv.svg)](https://pkg.go.dev/github.com/drewgonzales360/goenv)

`goenv` is an small, simple binary that executes the [install instructions](https://go.dev/doc/install) on the Go website and manages several Go versions. `goenv` downloads and extracts go to `/usr/local/goenv/<VERSION>` and adds a symlink from `/usr/local/go -> /usr/local/goenv/<VERSION>` (by default). It was loosely inspired by [Dave Cheney's blog post](https://dave.cheney.net/2014/04/20/how-to-install-multiple-versions-of-go).

## Usage

Calling `goenv` without any arguments will print out a helpful block of text, but here are a few useful examples. Note that installing 1.14 will install 1.14, even if 1.14.5 is the latest patch version.

```bash
# Helpful list of commands
❯ goenv
Usage:
   [command]

Available Commands:
  config      Print out the current config
  install     Install a Go version. Usually in the form 1.18, 1.9, 1.17.8.
  list        List all installed available Go versions.
  uninstall   Uninstall a Go version.
  use         Switch the current Go version to use whichever version in specified and installed.

Flags:
  -h, --help      help for this command
  -v, --version   version for this command

Use " [command] --help" for more information about a command.

# Install and use a go version
$ goenv install 1.17
✅ Downloaded and validated Go 1.17
✅ Extracted package
😎 Using go version go1.17 darwin/amd64

# Use an installed version. This will fail if you don't have it installed.
$ goenv use 1.19
😎 Using go version go1.19 darwin/amd64

# Removes an installed version and switches to another one if available
$ goenv rm 1.16
😎 Using go version go1.17.8 linux/amd64
😎 Uninstalled Go 1.16

# Lists installed versions
$ goenv list
Installed Versions:
1.17: 1.17
1.19: 1.19

# Shows the go installation location.
❯ goenv config
GOENV_ROOT_DIR:    /Users/me/.local/go    (set by environment variable)
GOENV_INSTALL_DIR: /Users/me/.local/goenv (set by environment variable)
```

## Install

To install goenv, follow the steps below. Older releases are in the [Releases page](https://github.com/drewgonzales360/goenv/releases).

```bash
# Step 1: Download Goenv for your unix based system and add it to /usr/local/bin
curl -sSL "https://github.com/drewgonzales360/goenv/releases/download/XXLatestXX/goenv-$(uname | tr '[:upper:]' '[:lower:]')-$(uname -m | sed 's/x86_64/amd64/')-XXLatestXX.tar.gz" | sudo tar -xzv -C /usr/local/bin

# Step 2: Add /usr/local/go/bin (or $GOENV_ROOT_DIR/bin) to PATH
export PATH=/usr/local/go/bin:PATH
```

Install this binary _without_ `go install` so that it is managed independent of Go.

### Configuration

| Environment Variable  | Default             | Explanation |
| -                     | -                   | - |
| GOENV_ROOT_DIR        | "/usr/local/go"     | This is the defaults to the default Go root. This directory /bin should be in your path and links to the GOENV_INSTALL_DIR |
| GOENV_INSTALL_DIR     | "/usr/local/goenv"  | Directory where your various Go installations will be installed |

The default `GOROOT` usually requires root access. You can avoid it by setting the configuration variables above. Adding the following to your .dotfiles will change the locations of the `GOENV_ROOT_DIR` and `GOENV_INSTALL_DIR`.

```shell
# goenv configuration, mkdir $HOME/.local if it doesn't exist
export GOENV_INSTALL_DIR="$HOME/.local/goenv"
export GOENV_ROOT_DIR="$HOME/.local/go"
export PATH="$GOENV_ROOT_DIR/bin:$PATH"
```

If your VScode editor throws you weird errors on start up, add this to your vscode settings.json. This happens because of the non-default `GOROOT`.
```json
{
    "go.gopath": "/your/gopath",
    "go.goroot": "/path/to/.local/go", # or whatever $GOENV_ROOT_DIR is set to
}
```

## How This Works

`goenv` does the following by default:
  - Checks for available versions of Go. This avoids code changes to `goenv` every time a new version is released.
  - Downloads the tarball for the corresponding version a user provides. Makes a best effort attempt to check the shasums
  - Extracts the tarball to `/usr/local/goenv/${VERSION}`. For example, `goenv install 1.17.6` will create `/usr/local/goenv/1.17.6`
  - Creates a symlink from `/usr/local/go/` to `/usr/local/goenv/1.17.6/`.

The install directory `/usr/local/goenv` and root directory `/usr/local/go` is configurable through environment variables.

## Other Implementations

There are a few other implementations of this that have more features, this project believes in doing the most simple task to manage Go versions. Other implementations can do fancier things like check for the correct Go version for a module and use the corresponding version when called, but they require intercepting the call to Go and passing the command to the right version. Goenv avoids being in the "hot path" and trusts users to know when to switch versions.

The recommendation on the Go website is to use your first install of Go to install _other_ versions of Go. Then you'd call other versions of Go like `go1.17.8 build`. This provides a consistent experience when switching versions.

Other tools:
- [syndbg/goenv](https://github.com/syndbg/goenv)
- [Spacewalkio/Goenv](https://github.com/Spacewalkio/Goenv)
- [moovweb/gvm](https://github.com/moovweb/gvm)
