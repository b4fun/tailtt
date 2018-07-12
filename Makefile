all:
	@echo 'available commands:'
	@echo
	@echo '	build	Build binary'

.PHONY: build
build:
	@go build -o bin/tailtt ./cmd/...
