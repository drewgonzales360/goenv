BUILD_METADATA=+$(shell git rev-parse --short HEAD)
PRERELEASE=
SEMVER=0.0.1

build:
	@go build -ldflags="-X 'github.com/drewgonzales360/goenv/version.Semver=${SEMVER}${PRERELEASE}${BUILD_METADATA}'"

image: build
	@docker build -t goenv .

run-image:
	@docker run --rm -it goenv bash -c 'goenv install 1.18; go version'

tty:
	@docker run --rm -it goenv bash
