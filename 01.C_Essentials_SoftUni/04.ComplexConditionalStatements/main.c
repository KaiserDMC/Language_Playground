#include <stdio.h>
#include <string.h>

void PersonalTitle();

void SmallShop();

void FruitOrVegetable();

void InvalidNumber();

void FruitShop();

void TradeCommissions();

void SkiTrip();

void OnTimeForTheExam();

int main() {
    OnTimeForTheExam();
    return 0;
}

void PersonalTitle() {
    char gender;
    double age;

    scanf("%lf", &age);
    scanf(" %c", &gender);

    switch (gender) {
        case 'm':
            if (age >= 16) {
                printf("Mr.");
            } else {
                printf("Master");
            }
            break;
        case 'f':
            if (age >= 16) {
                printf("Ms.");
            } else {
                printf("Miss");
            }
            break;
    }
}

void SmallShop() {
    char city[20];
    char product[20];
    double quantity;
    double price = 0.0;

    scanf("%s", product);
    scanf("%s", city);
    scanf("%lf", &quantity);

    if (strcmp(city, "Sofia") == 0) {
        if (strcmp(product, "coffee") == 0) {
            price = 0.50;
        } else if (strcmp(product, "water") == 0) {
            price = 0.80;
        } else if (strcmp(product, "beer") == 0) {
            price = 1.20;
        } else if (strcmp(product, "sweets") == 0) {
            price = 1.45;
        } else if (strcmp(product, "peanuts") == 0) {
            price = 1.60;
        }
    } else if (strcmp(city, "Plovdiv") == 0) {
        if (strcmp(product, "coffee") == 0) {
            price = 0.40;
        } else if (strcmp(product, "water") == 0) {
            price = 0.70;
        } else if (strcmp(product, "beer") == 0) {
            price = 1.15;
        } else if (strcmp(product, "sweets") == 0) {
            price = 1.30;
        } else if (strcmp(product, "peanuts") == 0) {
            price = 1.50;
        }
    } else if (strcmp(city, "Varna") == 0) {
        if (strcmp(product, "coffee") == 0) {
            price = 0.45;
        } else if (strcmp(product, "water") == 0) {
            price = 0.70;
        } else if (strcmp(product, "beer") == 0) {
            price = 1.10;
        } else if (strcmp(product, "sweets") == 0) {
            price = 1.35;
        } else if (strcmp(product, "peanuts") == 0) {
            price = 1.55;
        }
    }

    printf("%g", price * quantity);
}

void FruitOrVegetable() {
    char product[20];

    scanf("%s", product);

    if (strcmp(product, "banana") == 0 || strcmp(product, "apple") == 0 || strcmp(product, "kiwi") == 0 ||
        strcmp(product, "cherry") == 0 || strcmp(product, "lemon") == 0 || strcmp(product, "grapes") == 0) {
        printf("fruit");
    } else if (strcmp(product, "tomato") == 0 || strcmp(product, "cucumber") == 0 || strcmp(product, "pepper") == 0 ||
               strcmp(product, "carrot") == 0) {
        printf("vegetable");
    } else {
        printf("unknown");
    }
}

void InvalidNumber() {
    int number;

    scanf("%d", &number);

    if (number < 100 || number > 200) {
        if (number != 0) {
            printf("invalid");
        }
    }
}

void FruitShop() {
    char fruit[20];
    char day[20];
    double quantity, price = 0.0;

    scanf("%s %s %lf", fruit, day, &quantity);


    if (strcmp(day, "Saturday") == 0 || strcmp(day, "Sunday") == 0) {
        if (strcmp(fruit, "banana") == 0) {
            price = 2.70;
        } else if (strcmp(fruit, "apple") == 0) {
            price = 1.25;
        } else if (strcmp(fruit, "orange") == 0) {
            price = 0.90;
        } else if (strcmp(fruit, "grapefruit") == 0) {
            price = 1.60;
        } else if (strcmp(fruit, "kiwi") == 0) {
            price = 3.00;
        } else if (strcmp(fruit, "pineapple") == 0) {
            price = 5.60;
        } else if (strcmp(fruit, "grapes") == 0) {
            price = 4.20;
        } else {
            printf("error");
            return;
        }
    } else if (strcmp(day, "Monday") == 0 || strcmp(day, "Tuesday") == 0 || strcmp(day, "Wednesday") == 0 ||
               strcmp(day, "Thursday") == 0 || strcmp(day, "Friday") == 0) {
        if (strcmp(fruit, "banana") == 0) {
            price = 2.50;
        } else if (strcmp(fruit, "apple") == 0) {
            price = 1.20;
        } else if (strcmp(fruit, "orange") == 0) {
            price = 0.85;
        } else if (strcmp(fruit, "grapefruit") == 0) {
            price = 1.45;
        } else if (strcmp(fruit, "kiwi") == 0) {
            price = 2.70;
        } else if (strcmp(fruit, "pineapple") == 0) {
            price = 5.50;
        } else if (strcmp(fruit, "grapes") == 0) {
            price = 3.85;
        } else {
            printf("error");
            return;
        }
    } else {
        printf("error");
        return;
    }

    printf("%.2lf", price * quantity);
}

void TradeCommissions() {
    char city[20];
    double sales;

    scanf("%s", city);
    scanf("%lf", &sales);

    if (strcmp(city, "Sofia") == 0) {
        if (sales >= 0 && sales <= 500) {
            printf("%.2lf", sales * 0.05);
        } else if (sales > 500 && sales <= 1000) {
            printf("%.2lf", sales * 0.07);
        } else if (sales > 1000 && sales <= 10000) {
            printf("%.2lf", sales * 0.08);
        } else if (sales > 10000) {
            printf("%.2lf", sales * 0.12);
        } else {
            printf("error");
        }
    } else if (strcmp(city, "Varna") == 0) {
        if (sales >= 0 && sales <= 500) {
            printf("%.2lf", sales * 0.045);
        } else if (sales > 500 && sales <= 1000) {
            printf("%.2lf", sales * 0.075);
        } else if (sales > 1000 && sales <= 10000) {
            printf("%.2lf", sales * 0.10);
        } else if (sales > 10000) {
            printf("%.2lf", sales * 0.13);
        } else {
            printf("error");
        }
    } else if (strcmp(city, "Plovdiv") == 0) {
        if (sales >= 0 && sales <= 500) {
            printf("%.2lf", sales * 0.055);
        } else if (sales > 500 && sales <= 1000) {
            printf("%.2lf", sales * 0.08);
        } else if (sales > 1000 && sales <= 10000) {
            printf("%.2lf", sales * 0.12);
        } else if (sales > 10000) {
            printf("%.2lf", sales * 0.145);
        } else {
            printf("error");
        }
    } else {
        printf("error");
    }
}

void SkiTrip() {
    int days, nights;
    char roomType[40];
    char evaluation[20];
    double price = 0.0, total = 0.0, discount = 0.0;

    scanf("%d", &days);
    getchar();  // Consume the newline character left in the buffer
    fgets(roomType, sizeof(roomType), stdin);
    roomType[strlen(roomType) - 1] = '\0';  // Remove the trailing newline
    scanf("%s", evaluation);

    nights = days - 1;

    if (strcmp(roomType, "room for one person") == 0) {
        price = 18;
    } else if (strcmp(roomType, "apartment") == 0) {
        price = 25;

        if (nights >= 10 && nights <= 15) {
            discount = 0.35;
        } else if (nights > 15) {
            discount = 0.5;
        } else {
            discount = 0.3;
        }
    } else if (strcmp(roomType, "president apartment") == 0) {
        price = 35;

        if (nights >= 10 && nights <= 15) {
            discount = 0.15;
        } else if (nights > 15) {
            discount = 0.2;
        } else {
            discount = 0.1;
        }
    }

    total = price * nights;
    if (discount != 0) {
        total *= (1 - discount);
    }

    if (strcmp(evaluation, "positive") == 0) {
        total *= 1.25;
    } else {
        total *= 0.9;
    }

    printf("%.2lf", total);
}

void OnTimeForTheExam(){
    int examHour, examMinutes, arrivalHour, arrivalMinutes;
    
    scanf("%d %d %d %d", &examHour, &examMinutes, &arrivalHour, &arrivalMinutes);
    
    int examTime = examHour * 60 + examMinutes;
    int arrivalTime = arrivalHour * 60 + arrivalMinutes;
    
    if (arrivalTime > examTime) {
        printf("Late\n");
        int difference = arrivalTime - examTime;
        if (difference < 60) {
            printf("%d minutes after the start", difference);
        } else {
            int hours = difference / 60;
            int minutes = difference % 60;
            printf("%d:%02d hours after the start", hours, minutes);
        }
    } else if (arrivalTime == examTime || examTime - arrivalTime <= 30) {
        printf("On time\n");
        if (examTime - arrivalTime != 0) {
            printf("%d minutes before the start", examTime - arrivalTime);
        }
    } else {
        printf("Early\n");
        int difference = examTime - arrivalTime;
        if (difference < 60) {
            printf("%d minutes before the start", difference);
        } else {
            int hours = difference / 60;
            int minutes = difference % 60;
            printf("%d:%02d hours before the start", hours, minutes);
        }
    }
}