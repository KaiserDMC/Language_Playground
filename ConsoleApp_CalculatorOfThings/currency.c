#include "currency.h"
#include <stdio.h>

void display_submenu_currency(){
    printf("+++++++++++++++++++++++++++++++++++++++++++++++++++\n");
    printf("                      CURRENCY                     \n");
    printf("+++++++++++++++++++++++++++++++++++++++++++++++++++\n");
}

double calculate_currency(double amount, double rate){
    return amount / rate;
}

void currency_converter(){
    char origin[4];
    char convert[4];
    double amount, rate;

    display_submenu_currency();

    printf("Enter origin currency:\n");
    scanf("%s", origin);
    printf("Enter currency to convert to:\n");
    scanf("%s", convert);

    printf("Enter amount to convert:\n");
    scanf("%lf", &amount);
    printf("Enter current exchange rate:\n");
    scanf("%lf", &rate);

    printf("Converted amount of %s into %s: %.2lf\n", origin, convert, calculate_currency(amount, rate));
}