.PHONY: generate

BUILD_DIR   := build
BIN_NAME    := fhir-facade
BIN_VERSION := $(shell git describe --tags --always --long --dirty)
CGO_ENABLED := 0
GOARCH      := amd64


build:
	GOOS=linux   go build -C ./cmd/facade -ldflags "-X main.Version=${BIN_VERSION}" -o ../../${BUILD_DIR}/${BIN_NAME}-linux-${GOARCH}
	GOOS=darwin  go build -C ./cmd/facade -ldflags "-X main.Version=${BIN_VERSION}" -o ../../${BUILD_DIR}/${BIN_NAME}-darwin-${GOARCH}
	GOOS=windows go build -C ./cmd/facade -ldflags "-X main.Version=${BIN_VERSION}" -o ../../${BUILD_DIR}/${BIN_NAME}-windows-${GOARCH}.exe

clean:
	go clean
	rm -rf ./${BUILD_DIR}

generate:
	go run ./generate


NEW_VERSION = $(shell grep -om1 'v\d.\d.\d' CHANGELOG.md)

release:
	git tag ${NEW_VERSION}
	git push
	git push --tags
