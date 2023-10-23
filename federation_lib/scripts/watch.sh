#! /usr/bin/env sh
watchexec -w Src -w Tests -e h,cpp 'clear && scripts/run.sh'
