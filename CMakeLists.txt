cmake_minimum_required(VERSION 2.8.11)
project(SInGe)
#set(CMAKE_VERBOSE_MAKEFILE ON)

cmake_policy(SET CMP0042 NEW)

set(CMAKE_CXX_FLAGS "${CMAKE_CXX_FLAGS} -std=c++11 -Wall -Wextra -fPIC")
set(LIBRARY_OUTPUT_PATH ${PROJECT_BINARY_DIR}/lib)
set(EXECUTABLE_OUTPUT_PATH ${PROJECT_BINARY_DIR}/bin)

enable_testing()
add_subdirectory(src/)
add_subdirectory(golang/)

