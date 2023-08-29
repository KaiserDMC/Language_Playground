#ifndef CONSOLEAPP_CALCULATOROFTHINGS_TEMPERATURE_H
#define CONSOLEAPP_CALCULATOROFTHINGS_TEMPERATURE_H

void display_submenu_temperature();
double calculate_temperature_fahrenheit(double celsius);
double calculate_temperature_celsius(double fahrenheit);
double calculate_temperature(double temperature, char scale_from, char scale_to);
void temperature_converter();

#endif //CONSOLEAPP_CALCULATOROFTHINGS_TEMPERATURE_H
