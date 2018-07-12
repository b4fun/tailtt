all:
	@echo 'available commands:'
	@echo '	build	Build binary'

.PHONY: build
build:
	@./script/build.sh
