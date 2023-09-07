#include <stdio.h>

void display_submenu_bmr(){
    printf("+++++++++++++++++++++++++++++++++++++++++++++++++++\n");
    printf("                    BMR CALCULATOR                 \n");
    printf("+++++++++++++++++++++++++++++++++++++++++++++++++++\n");
}

double calculate_bmr_female(double weight, double height, double age){
    return ((10 * weight) + (6.25 * height) - (5 * age) - 161);
}

double calculate_bmr_male(double weight, double height, double age){
    return ((10 * weight) + (6.25 * height) - (5 * age) + 5);
}

void bmr_calculator(){
    char gender[2];
    double weight, height, age;

    display_submenu_bmr();

    printf("Enter which gender your are:\n");
    printf("(Example, enter m or M for Male and f or F for Female)\n");
    scanf("%s", gender);

    printf("Enter your weight:\n");
    scanf("%lf", &weight);

    printf("Enter your height:\n");
    scanf("%lf", &height);

    printf("Enter your age:\n");
    scanf("%lf", &age);

    if (gender[0] == 'm' || gender[0] == 'M'){
        printf("Calculated BMR for Male: %.2lf\n", calculate_bmr_male(weight, height, age));
    } else if (gender[0] == 'f' || gender[0] == 'F'){
        printf("Calculated BMR for Female: %.2lf\n", calculate_bmr_female(weight, height, age));
    } else {
        printf("Invalid gender!\n");
    }
}