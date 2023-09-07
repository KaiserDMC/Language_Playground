#include <stdio.h>
#include <string.h>
#include <math.h>

void ExcellentResult();

void GreaterNumber();

void EvenOrOdd();

void NumbersOneToNine();

void ThreeEqualNumbers();

void Numbers100To200();

void PasswordGuess();

void AreaOfFigures();

void DayOfWeek();

void AnimalType();

void ToyShop();

void Scholarship();

int main() {
    Scholarship();
    return 0;
}

void ExcellentResult() {
    float grade;
    scanf("%f", &grade);

    if (grade >= 5.50) {
        printf("Excellent!");
    }
}

void GreaterNumber() {
    int numOne, numTwo;

    scanf("%d", &numOne);
    scanf("%d", &numTwo);

    if (numOne > numTwo) {
        printf("%d", numOne);
    } else {
        printf("%d", numTwo);
    }
}

void EvenOrOdd() {
    int num;
    scanf("%d", &num);

    if (num % 2 == 0) {
        printf("even");
    } else {
        printf("odd");
    }
}

void NumbersOneToNine() {
    int num;
    scanf("%d", &num);
    char str[20];

    switch (num) {
        case 1:
            strcpy(str, "one");
            break;
        case 2:
            strcpy(str, "two");
            break;
        case 3:
            strcpy(str, "three");
            break;
        case 4:
            strcpy(str, "four");
            break;
        case 5:
            strcpy(str, "five");
            break;
        case 6:
            strcpy(str, "six");
            break;
        case 7:
            strcpy(str, "seven");
            break;
        case 8:
            strcpy(str, "eight");
            break;
        case 9:
            strcpy(str, "nine");
            break;
        default:
            strcpy(str, "number too big");
            break;
    }

    printf("%s", str);
}

void ThreeEqualNumbers() {
    int numOne, numTwo, numThree;
    scanf("%d", &numOne);
    scanf("%d", &numTwo);
    scanf("%d", &numThree);

    if (numOne == numTwo && numTwo == numThree) {
        printf("yes");
    } else {
        printf("no");
    }
}

void Numbers100To200() {
    int num;
    scanf("%d", &num);

    if (num < 100) {
        printf("Less than 100");
    } else if (num >= 100 && num <= 200) {
        printf("Between 100 and 200");
    } else {
        printf("Greater than 200");
    }
}

void PasswordGuess() {
    char password[40];
    scanf("%s", password);

    if (strcmp(password, "s3cr3t!P@ssw0rd") == 0) {
        printf("Welcome");
    } else {
        printf("Wrong password!");
    }
}

void AreaOfFigures() {
    char figure[20];
    scanf("%s", figure);

    if (strcmp(figure, "square") == 0) {
        float side;
        scanf("%f", &side);
        printf("%.3f", side * side);
    } else if (strcmp(figure, "rectangle") == 0) {
        float sideOne, sideTwo;
        scanf("%f", &sideOne);
        scanf("%f", &sideTwo);
        printf("%.3f", sideOne * sideTwo);
    } else if (strcmp(figure, "circle") == 0) {
        float radius;
        scanf("%f", &radius);
        printf("%.3f", M_PI * radius * radius);
    } else if (strcmp(figure, "triangle") == 0) {
        float side, height;
        scanf("%f", &side);
        scanf("%f", &height);
        printf("%.3f", side * height / 2);
    }
}

void DayOfWeek() {
    int day;
    scanf("%d", &day);

    switch (day) {
        case 1:
            printf("Monday");
            break;
        case 2:
            printf("Tuesday");
            break;
        case 3:
            printf("Wednesday");
            break;
        case 4:
            printf("Thursday");
            break;
        case 5:
            printf("Friday");
            break;
        case 6:
            printf("Saturday");
            break;
        case 7:
            printf("Sunday");
            break;
        default:
            printf("Error");
            break;
    }
}

void AnimalType() {
    char animal[20];
    scanf("%s", animal);

    if (strcmp(animal, "dog") == 0) {
        printf("mammal");
    } else if (strcmp(animal, "crocodile") == 0 || strcmp(animal, "tortoise") == 0 || strcmp(animal, "snake") == 0) {
        printf("reptile");
    } else {
        printf("unknown");
    }
}

void ToyShop() {
    double vacationPrice;
    int puzzleCount, dollCount, teddyBearCount, minionsCount, trucksCount;

    scanf("%lf", &vacationPrice);
    scanf("%d", &puzzleCount);
    scanf("%d", &dollCount);
    scanf("%d", &teddyBearCount);
    scanf("%d", &minionsCount);
    scanf("%d", &trucksCount);

    double totalProfit =
            puzzleCount * 2.6 + dollCount * 3 + teddyBearCount * 4.1 + minionsCount * 8.2 + trucksCount * 2;

    int totalCount = puzzleCount + dollCount + teddyBearCount + minionsCount + trucksCount;

    if (totalCount >= 50) {
        totalProfit *= 0.75;
    }

    totalProfit *= 0.9;

    if (totalProfit >= vacationPrice) {
        printf("Yes! %.2f lv left.", totalProfit - vacationPrice);
    } else {
        printf("Not enough money! %.2f lv needed.", vacationPrice - totalProfit);
    }
}

void Scholarship() {
    double income, averageGrade, minSalary;

    scanf("%lf", &income);
    scanf("%lf", &averageGrade);
    scanf("%lf", &minSalary);

    double socialScholarship = floor(minSalary * 0.35);
    double excellentScholarship = floor(averageGrade * 25);

    if (averageGrade >= 5.5) {
        if (excellentScholarship < socialScholarship && income < minSalary) {
            printf("You get a Social scholarship %.0f BGN", socialScholarship);
        } else {
            printf("You get a scholarship for excellent results %.0f BGN", excellentScholarship);
        }
    } else if (income < minSalary && averageGrade > 4.5) {

        printf("You get a Social scholarship %.0f BGN", socialScholarship);

    } else {
        printf("You cannot get a scholarship!");
    }
}