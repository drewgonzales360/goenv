.EXPORT_ALL_VARIABLES:
BUILD_METADATA=+$(shell git rev-parse --short HEAD)
PRERELEASE=
SEMVER=v0.0.3
VERSION=${SEMVER}${PRERELEASE}${BUILD_METADATA}

# Builds target for whatever OS this is called from.
build:
	go build -ldflags="-X 'github.com/drewgonzales360/goenv/version.Semver=${VERSION}'"

install:
	mv goenv /usr/local/bin

build-linux:
	GOOS=linux go build -ldflags="-X 'github.com/drewgonzales360/goenv/version.Semver=${VERSION}'"

build-darwin:
	GOOS=darwin go build -ldflags="-X 'github.com/drewgonzales360/goenv/version.Semver=${VERSION}'"

image: build-linux
	docker build -t goenv .

# Runs a script to test basic, happy-path functionality inside the container
test: image
	docker run --rm -it -e GOENV_LOG=DEBUG --entrypoint bash goenv goenv-test

# Opens up a container to play around with goenv. Installing, removing, and switching go versions
# is much safer in the container than it is on your local machine. It is short for interactive.
it: image
	docker run --rm -it -e GOENV_LOG=DEBUG goenv

# This creates a github release, but requires the caller to be properly authenticated
# Only I, drewgonzales360, can create releases right now.
release:
	@if ! [ -d tmp ]; then mkdir tmp; fi
	@GOOS=linux make build
	@tar -czf tmp/goenv-linux-amd64-${SEMVER}.tar.gz ./goenv
	@GOOS=darwin make build
	@tar -czf tmp/goenv-darwin-amd64-${SEMVER}.tar.gz ./goenv
	@git tag ${SEMVER}
	gh release create --notes "Release ${VERSION}" --target main ${SEMVER} tmp/goenv-*-amd64-${SEMVER}.tar.gz

readme:
	envsubst < templates/README.md > README.md

# Turns on some hooks to check format and build status before commiting/pushing. Optional, but helpful.
githooks:
	git config --local core.hooksPath .githooks/

clean:
	@if [ -d tmp ]; then rm -rf tmp; fi
	@if [ -f goenv ]; then rm goenv; fi
