##common
cmake_minimum_required(VERSION 2.5)
SET(CMAKE_CXX_FLAGS_DEBUG "$ENV{CXXFLAGS} -O0  -g -pipe -W -Wall -fPIC")
SET(CMAKE_CXX_FLAGS_RELEASE "$ENV{CXXFLAGS} -O3 -g -pipe -W -Wall -fPIC")
set(CMAKE_BUILD_TYPE Debug)


##project
project(collector)
SET(EXECUTABLE_OUTPUT_PATH ${PROJECT_BINARY_DIR}/bin)
set(SRC_LIST collector/regex.example.cpp util/util.cpp)
include_directories(util  lib/pcre-install/include)

## dynamic linker
#link_directories(  
#         ${PROJECT_SOURCE_DIR}/lib/pcre-install/lib) 

add_executable(collector  ${SRC_LIST})
target_link_libraries(collector ${CMAKE_SOURCE_DIR}/lib/pcre-install/lib/libpcre.a)
