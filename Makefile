.PHONY: build build-darwin-amd64 build-linux-amd64 build-checksums release

#
# Build
#

build: build-clean build-darwin-amd64 build-linux-amd64 build-checksums

build-clean:
	rm -f ./build/*

build-darwin-amd64:
	GOOS=darwin GOARCH=amd64 go build -o ./build/lokitool-darwin-amd64 ./cmd/lokitool

build-linux-amd64:
	GOOS=linux GOARCH=amd64 go build -o ./build/lokitool-linux-amd64 ./cmd/lokitool

build-checksums:
	for ARTIFACT in ./build/*; do \
		SHA256=$$(shasum -a 256 $${ARTIFACT} | awk '{print $$1 }'); \
		echo "$${SHA256}" > $${ARTIFACT}-sha256; \
	done

#
# Release
#

ARTIFACTS=$(shell ls ./build/*)
ARTIFACTS_ATTACH=$(addprefix --attach=, $(ARTIFACTS))

release: build
	hub release create --draft $(ARTIFACTS_ATTACH) v$(VERSION)
