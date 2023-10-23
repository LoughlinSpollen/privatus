#! /usr/bin/env sh
scripts/select-compiler.sh COMPILER=LLVM

mkdir -p build/lib
cd build/lib

cmake -DCMAKE_BUILD_TYPE=Debug -Wall -G "Ninja" ../..
cmake --build . -- -j14 || exit $?

cd ../..
exit
