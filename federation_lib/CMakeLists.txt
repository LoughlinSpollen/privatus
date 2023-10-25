cmake_minimum_required(VERSION 3.1)

set(CMAKE_CXX_STANDARD 17)
set(CMAKE_CXX_STANDARD_REQUIRED ON)
set(CMAKE_CXX_EXTENSIONS OFF)

if(NOT CMAKE_BUILD_TYPE)
  set(CMAKE_BUILD_TYPE Release)
endif()

set(CMAKE_CXX_FLAGS "-O3")
set(CMAKE_CXX_FLAGS_RELEASE "-O3")

message(STATUS "1 -------------------------------------" )

# If you are using a different py verions: 
# find site-packages  python -c 'import site; print(site.getsitepackages())'
set(PYTHON_LIBRARY_DIR "~/.pyenv/shims/versions/3.7.17/lib/python3.7/site-packages")
set(PYTHON_EXECUTABLE "~/.pyenv/shims/shims/python")

set(PYTHON_INCLUDE_DIR "~/.pyenv/shims/versions/3.7.17/include/python3.7m/")
set(PYTHON_LIBRARY "~/.pyenv/shims/versions/3.7.17/lib/python3.7/")

set(PYTHON_MODULE_EXTENSION ".so" CACHE INTERNAL "Cross python lib extension")

# pybind11 headers
set(pybind11_DIR "usr/local/include")

# Disable pybind11 python search mechanism
set(PYTHONLIBS_FOUND TRUE CACHE INTERNAL "")

project(privatus)

include_directories("${CMAKE_SOURCE_DIR}/cpp/include/domain/model" "${CMAKE_SOURCE_DIR}/cpp/include/infra/network" "${CMAKE_SOURCE_DIR}/cpp/include/usecase")
include_directories("${CMAKE_SOURCE_DIR}/python")

include_directories("${PYTHON_INCLUDE_DIR}")

file (GLOB SOURCE_FILES "cpp/src/domain/model/*.cpp" "cpp/src/infra/network/*.cpp" "cpp/src/usecase/*.cpp")
file (GLOB HEADER_FILES "cpp/include/domain/model/*.hpp" "cpp/include/infra/network/*.hpp" "/cpp/include/usecase/*.hpp")
file (GLOB PYTHON_FILES "python/*.cpp" "python/*.hpp")

# Set up such that XCode organizes the files
source_group(TREE ${CMAKE_CURRENT_SOURCE_DIR} FILES ${SOURCE_FILES} ${HEADER_FILES} ${PYTHON_FILES} )

include(pybind11.cmake)
pybind11_add_module(privatus 
	${SOURCE_FILES}
	${HEADER_FILES}
	${PYTHON_FILES}
)

target_link_libraries(privatus PUBLIC)

message(STATUS "2 -------------------------------------" )


install(TARGETS privatus
  COMPONENT python
  LIBRARY DESTINATION "${PYTHON_LIBRARY_DIR}"
)

message(STATUS "3 -------------------------------------" )
