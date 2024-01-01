cmake_minimum_required(VERSION 3.26)
project(ConsoleApp_CalculatorOfThings C)

set(CMAKE_C_STANDARD 17)

add_executable(ConsoleApp_CalculatorOfThings main.c
        menu.c
        currency.h
        currency.c
        temperature.h
        temperature.c
        temperature_table.h
        temperature_table.c
        bmi.h
        bmi.c
        bmr.h
        bmr.c
)
