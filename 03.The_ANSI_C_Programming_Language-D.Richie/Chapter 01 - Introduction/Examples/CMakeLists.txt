cmake_minimum_required(VERSION 3.28)
project(Examples C)

set(CMAKE_C_STANDARD 90)

add_executable(Examples main.c
        longest_input_special.c)
