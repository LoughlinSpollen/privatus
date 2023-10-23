#! /usr/bin/env zsh
cd Src
clang-format -i -style=LLVM ./**/*.cpp ./**/*.h
cd ..