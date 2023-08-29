#include <stdio.h>
#include <stdlib.h>
#include "currency.h"
#include "temperature.h"
#include "temperature_table.h"
#include "bmi.h"
#include "bmr.h"

void display_menu(){
    printf("+++++++++++++++++++++++++++++++++++++++++++++++++++\n");
    printf("                        MENU                        \n");
    printf("+++++++++++++++++++++++++++++++++++++++++++++++++++\n");
    printf(" Currency Converter                             :1 \n");
    printf(" Temperature Converter                          :2 \n");
    printf(" Temperature Table                              :3 \n");
    printf(" BMI Calculator                                 :4 \n");
    printf(" BMR Calculator                                 :5 \n");
    printf(" Exit                                           :0 \n");
    printf("+++++++++++++++++++++++++++++++++++++++++++++++++++\n");
    printf(" Which program would you like to run?: ");
}

void start(){
    char input[2];
    int choice = -1;

    while (choice != 0){
        display_menu();

        // Read user input as a string
        if (fgets(input, sizeof(input), stdin) == NULL) {
            // Handle input error, if any
            printf("Error reading input\n");
            exit(1);
        }

        // Convert the string to an integer using strtol
        choice = strtol(input, NULL, 10);

        switch (choice){
            case 1:
                currency_converter();
                break;
            case 2:
                temperature_converter();
                break;
            case 3:
                temperature_table();
                break;
            case 4:
                bmi_calculator();
                break;
            case 5:
                bmr_calculator();
                break;
            case 0:
                printf("Exiting...\n");
                return;
            default:
                printf("Invalid choice!\n");
                break;
        }

        choice = getchar();
    }
}