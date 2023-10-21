#!/usr/bin/env bash
watchexec -w scripts/build.sh -w ./*.proto --restart 'clear; ./scripts/build.sh'