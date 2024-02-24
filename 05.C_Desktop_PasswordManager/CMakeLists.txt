cmake_minimum_required(VERSION 3.26)
project(DesktopPasswordManager C)

set(CMAKE_CXX_STANDARD 17)

include_directories(.)

add_library(DesktopPasswordManager SHARED
        main.c)
