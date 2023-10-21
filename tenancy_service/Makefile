.PHONY: all test lint build run watch install-dev
.DEFAULT_GOAL := run
SHELL         := /bin/bash

build:
	@scripts/build.sh

lint:
	@scripts/lint.sh

run:
	@scripts/run.sh

watch:
	@scripts/watch.sh

test:
	@scripts/config.sh
	@scripts/test.sh

install-dev:
	@scripts/install-dev.sh
