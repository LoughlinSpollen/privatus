#! /usr/bin/env sh
scripts/select-compiler.sh  COMPILER=LLVM

rm -rf build/lib
mkdir -p build/lib
cd build/lib

cmake -g -DCMAKE_BUILD_TYPE=Release -Wall -G "Ninja" ../..
cmake --build . --target pyprivatus -- -j14 || exit $?

cd ../..
exit
