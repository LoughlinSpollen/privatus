.PHONY: all db-setup db-init db-start db-stop install-dev
.DEFAULT_GOAL := run
SHELL         := /bin/bash

db-setup:
	@scripts/db-stop.sh
	@scripts/db-delete.sh
	@scripts/db-start.sh
	@scripts/db-setup.sh
	@scripts/db-init.sh

db-init:
	@scripts/db-init.sh 

db-start:
	@scripts/db-start.sh

db-stop:
	@scripts/db-stop.sh

install-dev:
	@scripts/install-dev.sh
