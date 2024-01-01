#include <stdio.h>

void display_submenu_bmi(){
    printf("+++++++++++++++++++++++++++++++++++++++++++++++++++\n");
    printf("                      BMI                          \n");
    printf("+++++++++++++++++++++++++++++++++++++++++++++++++++\n");
}

double calculate_bmi(double weight, double height){
    return weight / (height * height);
}

double calculate_bmi_imperial(double weight, double height){
    return (weight * 703) / (height * height);
}

void bmi_calculator(){
    char scale[2];
    double weight, height;

    display_submenu_bmi();

    printf("Enter which units to use:\n");
    printf("(Example, enter m or M to use Metric and i or I to use Imperial)\n");
    scanf("%s", scale);

    printf("Enter your weight:\n");
    scanf("%lf", &weight);

    printf("Enter your height:\n");
    scanf("%lf", &height);

    if (scale[0] == 'm' || scale[0] == 'M'){
        printf("Calculated BMI in Metric: %.2lf\n", calculate_bmi(weight, height));
    } else if (scale[0] == 'i' || scale[0] == 'I'){
        printf("Calculated BMI in Imperial: %.2lf\n", calculate_bmi_imperial(weight, height));
    } else {
        printf("Invalid weight scale!\n");
    }
}