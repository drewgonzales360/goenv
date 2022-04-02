# Contributing

Pull Requests are welcome! I'd like to keep this project small, simple, and cute 🥳. The [Makefile](./Makefile) is heavily commented to help lower the barrier to contributions.

## How This Works

`goenv` does the following:
  - Downloads the tarball for the corresponding version a user provides
  - Extracts the tarball to `/usr/local/go/${VERSION}`. For example, `goenv install 1.17.6` will create `/usr/local/go/1.17.6`
  - Creates a symlink from `/usr/local/bin/go/` to `/usr/local/go/1.17.6/bin/go`. The same thing happens for `gofmt`.

This project resists adding more functionality in an attempt to be simple. The four available commands will be limited to that for the forseeable future.

## Other Implementations

There are a few other implementations of this that have more features, but I only needed the binaries. Other implementations can do fancier things like check for the correct Go version and use the corresponding version in runtime, but they require intercepting the call to Go and passing the command to the right version. I didn't want my code to be in the "hot path."

A few other implementations are also written in other languages. This one is written in Go 🥵.

The recommendation on the Go website is to use your first install of Go to install _other_ versions of Go. Then you'd call other versions of Go like `go1.17.8 build`. This provides a consistent experience.
