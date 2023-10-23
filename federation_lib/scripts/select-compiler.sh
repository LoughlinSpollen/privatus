#! /usr/bin/env bash

COMPILER=${COMPILER:-LLVM}
echo "Compiler: ${COMPILER}" 
if [ "${COMPILER}" == "LLVM" ]
then
  export C=/usr/bin/clang
  export CC=/usr/bin/clang
  export CXX=/usr/bin/clang++
else
  export C=/usr/local/bin/gcc-13
  export CC=/usr/local/bin/gcc-13
  export CXX=/usr/local/bin/g++-13
fi
