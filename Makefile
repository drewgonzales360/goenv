BUILD_METADATA=+$(shell git rev-parse --short HEAD)
PRERELEASE=
SEMVER=0.0.1

build:
	@go build -ldflags="-X 'github.com/drewgonzales360/goenv/version.Semver=${SEMVER}${PRERELEASE}${BUILD_METADATA}'"

build-linux:
	@GOOS=linux go build -ldflags="-X 'github.com/drewgonzales360/goenv/version.Semver=${SEMVER}${PRERELEASE}${BUILD_METADATA}'"

image: build-linux
	@docker build -t goenv .

test: image
	@docker run --rm -it --entrypoint bash goenv goenv-test

ty:
	@docker run --rm -it goenv
