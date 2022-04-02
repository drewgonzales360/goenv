BUILD_METADATA=+$(shell git rev-parse --short HEAD)
PRERELEASE=
SEMVER=v0.0.2
VERSION=${SEMVER}${PRERELEASE}${BUILD_METADATA}

# Builds target for whatever OS this is called from.
build:
	@go build -ldflags="-X 'github.com/drewgonzales360/goenv/version.Semver=${VERSION}'"

# Builds it for linux
build-linux:
	@GOOS=linux go build -ldflags="-X 'github.com/drewgonzales360/goenv/version.Semver=${VERSION}'"

image: build-linux
	@docker build -t goenv .

# Runs a script to test basic, happy-path functionality inside the container
test: image
	@docker run --rm -it --entrypoint bash goenv goenv-test

# Opens up a container to play around with goenv. Installing, removing, and switching go versions
# is much safer in the container than it is on your local machine. It is short for interactive.
it:
	@docker run --rm -it goenv

# This creates a github release, but requires the caller to be properly authenticated
# Only I, drewgonzales360, can create releases right now.
release: build-linux
	@if ! [ -d tmp ]; then mkdir tmp; fi
	@tar -czf tmp/goenv-amd64-${SEMVER}.tar.gz ./goenv
	gh release create ${SEMVER} "tmp/goenv-amd64-${SEMVER}.tar.gz" --notes "Release ${VERSION}"

# Turns on some hooks to check format and build status before commiting/pushing. Optional, but helpful.
githooks:
	git config --local core.hooksPath .githooks/

clean:
	@if [ -d tmp ]; then rm -rf tmp; fi
	@if [ -f tmp ]; then rm goenv; fi
