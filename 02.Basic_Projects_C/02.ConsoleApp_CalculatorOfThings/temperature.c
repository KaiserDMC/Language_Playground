#include <stdio.h>

void display_submenu_temperature(){
    printf("+++++++++++++++++++++++++++++++++++++++++++++++++++\n");
    printf("                    TEMPERATURE                    \n");
    printf("+++++++++++++++++++++++++++++++++++++++++++++++++++\n");
}

double calculate_temperature_fahrenheit(double celsius){
    return (celsius * 9/5) + 32;
}

double calculate_temperature_celsius(double fahrenheit){
    return (fahrenheit - 32) * 5/9;
}

double calculate_temperature(double temperature, char scale_from, char scale_to) {
    if (scale_from == 'c' || scale_from == 'C') {
        if (scale_to == 'f' || scale_to == 'F') {
            return (temperature * 9/5) + 32; // Celsius to Fahrenheit
        }
    } else if (scale_from == 'f' || scale_from == 'F') {
        if (scale_to == 'c' || scale_to == 'C') {
            return (temperature - 32) * 5/9; // Fahrenheit to Celsius
        }
    }

    return 0.0 / 0.0; // NaN
}

void temperature_converter(){
    char scale[2];
    double temperature;

    display_submenu_temperature();

    printf("Enter temperature scale to convert FROM:\n");
    printf("(Example, enter f or F to convert Fahrenheit to Celsius)\n");
    scanf("%s", scale);

    printf("Enter temperature to convert:\n");
    scanf("%lf", &temperature);

    if (scale[0] == 'c' || scale[0] == 'C'){
        printf("Converted %.2lf from Celsius to Fahrenheit: %.2lf\n", temperature, calculate_temperature_fahrenheit(temperature));
    } else if (scale[0] == 'f' || scale[0] == 'F'){
        printf("Converted %.2lf from Fahrenheit to Celsius: %.2lf\n", temperature, calculate_temperature_celsius(temperature));
    } else {
        printf("Invalid temperature scale!\n");
    }
}