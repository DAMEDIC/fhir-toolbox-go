.PHONY: generate

BUILD_DIR   := build

clean:
	go clean
	rm -rf ./${BUILD_DIR}

generate:
	go run ./internal/cmd/generate
