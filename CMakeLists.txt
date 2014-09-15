##common
cmake_minimum_required(VERSION 3.0)
SET(CMAKE_CXX_FLAGS_DEBUG "$ENV{CXXFLAGS} -O0 -Wall -g -ggdb -lpcre")
SET(CMAKE_CXX_FLAGS_RELEASE "$ENV{CXXFLAGS} -O3 -Wall")
set(CMAKE_BUILD_TYPE Debug)

##lib
set(PCRELIB lib/pcre-install/lib) 

##project
project(collector)
set(SRC_LIST collector/regex.example.cpp util/util.cpp)
#add_library(libpcre STATIC ${PCRELIB}/libpcre.a)
#add_library(libpcre SHARED ${PCRELIB}/libpcre.a)
#SET_TARGET_PROPERTIES(libpcre PROPERTIES LINKER_LANGUAGE C)
include_directories(util  lib/pcre-install/include)
#add_subdirectory( util )
#link_libraries(pcre ${PCRELIB}/libpcre.a)

add_executable(collector  ${SRC_LIST})
target_link_libraries(collector /Users/zealot/github/eyes/lib/pcre-install/lib/libpcre.a)
