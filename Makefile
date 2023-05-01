BUILD_METADATA=+$(shell git rev-parse --short HEAD)
PRERELEASE=
SEMVER=v0.2.0
VERSION=${SEMVER}${PRERELEASE}${BUILD_METADATA}
GOOS?=$(shell uname | tr '[:upper:]' '[:lower:]')
GOARCH?=$(shell uname -m | sed 's/x86_64/amd64/')

# Builds target for whatever OS this is called from.
# go tool nm ./goenv
build:
	@echo [info] building goenv for ${GOOS}-${GOARCH}
	@go build -ldflags="-X 'main.Semver=${VERSION}'"

tar: build
	@tar -czf tmp/goenv-${GOOS}-${GOARCH}-${SEMVER}.tar.gz ./goenv

install: build
	mv goenv /usr/local/bin

# Runs a script to test basic, happy-path functionality inside the container
test:
	@GOOS=linux go build -ldflags="-X 'main.Semver=${VERSION}'"
	@docker build -t goenv-ubuntu .
	docker run --rm -it -e GOENV_LOG=DEBUG --entrypoint /usr/local/bin/goenv-test goenv-ubuntu

# Opens up a container to play around with goenv. Installing, removing, and switching go versions
# is much safer in the container than it is on your local machine. It is short for interactive.
it:
	docker run --rm -it -e GOENV_LOG=DEBUG goenv-ubuntu

# This creates a github release, but requires the caller to be properly authenticated
# Only I, drewgonzales360, can create releases right now.
release:
	@if ! [ -d tmp ]; then mkdir tmp; fi
	@GOOS=linux GOARCH=amd64 make tar
	@GOOS=linux GOARCH=arm64 make tar
	@GOOS=darwin GOARCH=amd64 make tar
	@GOOS=darwin GOARCH=arm64 make tar
	@git tag ${SEMVER}
	gh release create --notes "Release ${VERSION}" --target main ${SEMVER} tmp/goenv-*-*-${SEMVER}.tar.gz
readme:
	@sed "s/XXLatestXX/${SEMVER}/g" < templates/README.md > README.md

# Turns on some hooks to check format and build status before commiting/pushing. Optional, but helpful.
githooks:
	git config --local core.hooksPath .githooks/

clean:
	@if [ -d tmp ]; then rm -rf tmp; fi
	@if [ -f goenv ]; then rm goenv; fi
