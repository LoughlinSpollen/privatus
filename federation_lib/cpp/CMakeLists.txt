cmake_minimum_required(VERSION 3.1)

set(CMAKE_CXX_STANDARD 17)
set(CMAKE_CXX_STANDARD_REQUIRED ON)
set(CMAKE_CXX_EXTENSIONS OFF)

project(privatus VERSION 0.1.0)

message(STATUS "4 -------------------------------------" )


include_directories(~/.pyenv/shims/versions/3.7.17/include/python3.7m/)

include_directories(/usr/local/include)
link_directories(/usr/local/lib)

AUX_SOURCE_DIRECTORY(src SRC_FILES)

set(PROJECT_LIB_SOURCE_DIR "src" )
set(PROJECT_LIB_INCLUDE_DIR "include")
set(PROJECT_MODEL_SOURCE_DIR "src/domain/model" )
set(PROJECT_MODEL_INCLUDE_DIR "include/domain/model")
set(PROJECT_INFRA_SOURCE_DIR "src/infra/network")
set(PROJECT_INFRA_INCLUDE_DIR "include/infra/network")
set(PROJECT_USECASE_SOURCE_DIR "src/usecase")
set(PROJECT_USECASE_INCLUDE_DIR "include/usecase")

set(SOURCE_FILES
    ${PROJECT_LIB_SOURCE_DIR}/PrivatusLib.hpp
    ${PROJECT_LIB_INCLUDE_DIR}/PrivatusLib.cpp

    ${PROJECT_MODEL_INCLUDE_DIR}/FederationConfig.hpp
    ${PROJECT_MODEL_SOURCE_DIR}/FederationConfig.cpp

    ${PROJECT_INFRA_INCLUDE_DIR}/FederationClient.hpp
    ${PROJECT_INFRA_SOURCE_DIR}/FederationClient.cpp

    ${PROJECT_USECASE_INCLUDE_DIR}/RegisterUsecase.hpp
    ${PROJECT_USECASE_SOURCE_DIR}/RegisterUsecase.cpp
)
message(STATUS "5 -------------------------------------" )


source_group(TREE ${CMAKE_CURRENT_SOURCE_DIR} FILES ${SOURCE_FILES})

add_library(privatus SHARED ${SOURCE_FILES})

target_include_directories(privatus PRIVATE ${PROJECT_LIB_INCLUDE_DIR} ${PROJECT_MODEL_INCLUDE_DIR} ${PROJECT_INFRA_INCLUDE_DIR} ${PROJECT_USECASE_INCLUDE_DIR})

install(TARGETS privatus DESTINATION lib)
install(FILES include/PrivatusLib DESTINATION include)
install(DIRECTORY ${PROJECT_MODEL_INCLUDE_DIR} DESTINATION include)
install(DIRECTORY ${PROJECT_INFRA_INCLUDE_DIR} DESTINATION include)
install(DIRECTORY ${PROJECT_USECASE_INCLUDE_DIR} DESTINATION include)


message(STATUS "6 -------------------------------------" )
