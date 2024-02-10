#include <stdio.h>
#include "temperature_table.h"
#include "temperature.h"

void display_submenu_temperature_table(){
    printf("+++++++++++++++++++++++++++++++++++++++++++++++++++\n");
    printf("                    TEMPERATURE TABLE              \n");
    printf("+++++++++++++++++++++++++++++++++++++++++++++++++++\n");
}

void temperature_table(){
    double temperature;
    int lower, upper, step;
    char scale_from, scale_to;

    display_submenu_temperature_table();

    printf("Enter lower limit:\n");
    scanf("%d", &lower);
    printf("Enter upper limit:\n");
    scanf("%d", &upper);
    printf("Enter step:\n");
    scanf("%d", &step);

    printf("From Scale: (c or f)\n");
    scanf(" %c", &scale_from);
    printf("To Scale: (c or f)\n");
    scanf(" %c", &scale_to);

    printf("From Scale: %c, To Scale: %c\n", scale_from, scale_to);
    printf("Original Scale Temperature Converted Temperature\n");
    for (int i = lower; i <= upper; i += step){
        temperature = i;
        double converted_temperature = calculate_temperature(temperature, scale_from, scale_to);
        printf("%.2lf %c %.2lf %c\n", temperature, scale_from, converted_temperature, scale_to);
    }
}