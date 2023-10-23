cmake_minimum_required(VERSION 3.1)

set(CMAKE_CXX_STANDARD 17)
set(CMAKE_CXX_STANDARD_REQUIRED ON)
set(CMAKE_CXX_EXTENSIONS OFF)

project(privatus VERSION 0.1.0)

include_directories(/usr/local/include)
include_directories(/usr/local/include)

AUX_SOURCE_DIRECTORY(src SRC_FILES)

set(PROJECT_MODEL_SOURCE_DIR "src/domain/model" )
set(PROJECT_MODEL_INCLUDE_DIR "include/domain/model")
set(PROJECT_INFRA_SOURCE_DIR "src/infra/network")
set(PROJECT_INFRA_INCLUDE_DIR "include/infra/network")
set(PROJECT_USECASE_SOURCE_DIR "cpp/src/usecase")
set(PROJECT_USECASE_INCLUDE_DIR "include/usecase")

set(SOURCE_FILES
    ${PROJECT_MODEL_INCLUDE_DIR}/FederationConfig.hpp
    ${PROJECT_MODEL_SOURCE_DIR}/FederationConfig.cpp

    ${PROJECT_INFRA_INCLUDE_DIR}/FederationService.hpp
    ${PROJECT_INFRA_SOURCE_DIR}/FederationService.cpp

    ${PROJECT_USECASE_INCLUDE_DIR}/RegisterUsecase.hpp
    ${PROJECT_USECASE_SOURCE_DIR}/RegisterUsecase.cpp
)

source_group(TREE ${CMAKE_CURRENT_SOURCE_DIR} FILES ${SOURCE_FILES})

add_library(privatus SHARED ${SOURCE_FILES})


target_include_directories(privatus PRIVATE ${PROJECT_MODEL_INCLUDE_DIR} ${PROJECT_INFRA_INCLUDE_DIR} ${PROJECT_USECASE_INCLUDE_DIR})

install(TARGETS privatus DESTINATION lib)
install(FILES include/privatus DESTINATION include)
install(DIRECTORY ${PROJECT_MODEL_INCLUDE_DIR} DESTINATION include)
install(DIRECTORY ${PROJECT_INFRA_INCLUDE_DIR} DESTINATION include)
install(DIRECTORY ${PROJECT_USECASE_INCLUDE_DIR} DESTINATION include)
