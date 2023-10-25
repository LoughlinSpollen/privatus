.PHONY: install-dev build-release build-debug clean watch test
.DEFAULT_GOAL := run

install-dev:
	@scripts/install-dev.sh

build-release:
	@scripts/clean.sh
	@scripts/build-release.sh

build-debug:
	@scripts/clean.sh
	@scripts/build-debug.sh

clean:
	@scripts/clean.sh

watch:
	@scripts/watch.sh

test:
	@scripts/test.sh
