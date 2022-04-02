BUILD_METADATA=+$(shell git rev-parse --short HEAD)
PRERELEASE=
SEMVER=0.0.1
VERSION=${SEMVER}${PRERELEASE}${BUILD_METADATA}

build:
	@go build -ldflags="-X 'github.com/drewgonzales360/goenv/version.Semver=${VERSION}'"

build-linux:
	@GOOS=linux go build -ldflags="-X 'github.com/drewgonzales360/goenv/version.Semver=${VERSION}'"

image: build-linux
	@docker build -t goenv .

test: image
	@docker run --rm -it --entrypoint bash goenv goenv-test

ty:
	@docker run --rm -it goenv

release: build-linux
	@mkdir tmp > /dev/null || true
	tar -cvzf tmp/goenv_amd64_${SEMVER}.tar.gz ./goenv
	gh release create ${SEMVER} "tmp/goenv_amd64_${SEMVER}.tar.gz#goenv-amd64-${SEMVER}.tar.gz" --notes "Release ${VERSION}"
