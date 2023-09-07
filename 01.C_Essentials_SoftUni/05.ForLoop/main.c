#include <stdio.h>
#include <limits.h>
#include <math.h>
#include <string.h>

void NumbersFrom1To100();

void LatinLetters();

void SumNumbers();

void NumberSequence();

void LeftAndRightSum();

void OddEvenSum();

void VowelsSum();

void CleverLily();

void DivideWithoutReminder();

int main() {
    DivideWithoutReminder();
    return 0;
}

void NumbersFrom1To100() {
    for (int i = 1; i <= 100; i++) {
        printf("%d\n", i);
    }
}

void LatinLetters() {
    for (char i = 'a'; i <= 'z'; i++) {
        printf("%c\n", i);
    }
}

void SumNumbers() {
    int initialNumber, number, sum = 0;
    scanf("%d", &initialNumber);

    for (int i = 1; i <= initialNumber; ++i) {
        scanf("%d", &number);
        sum += number;
    }

    printf("%d", sum);
}

void NumberSequence() {
    int initialNumber, number, maxNumber = INT_MIN, minNumber = INT_MAX;
    scanf("%d", &initialNumber);

    for (int i = 0; i < initialNumber; ++i) {

        scanf("%d", &number);

        if (number < minNumber) {
            minNumber = number;
        }

        if (number > maxNumber) {
            maxNumber = number;
        }
    }

    printf("Max number: %d\n", maxNumber);
    printf("Min number: %d", minNumber);
}

void LeftAndRightSum() {
    int initialNumber, number;
    double sumLeft = 0, sumRight = 0;
    scanf("%d", &initialNumber);

    for (int i = 1; i <= initialNumber; ++i) {
        scanf("%d", &number);
        sumLeft += number;
    }

    for (int i = 1; i <= initialNumber; ++i) {
        scanf("%d", &number);
        sumRight += number;
    }

    if (sumLeft == sumRight) {
        printf("Yes, sum = %lf", sumLeft);
    } else {
        printf("No, diff = %lf", fabs(sumLeft - sumRight));
    }
}

void OddEvenSum() {
    int initialNumber, number;
    double sumOdd = 0, sumEven = 0;
    scanf("%d", &initialNumber);

    for (int i = 1; i <= initialNumber; ++i) {
        scanf("%d", &number);

        if (i % 2 == 0) {
            sumEven += number;
        } else {
            sumOdd += number;
        }
    }

    if (sumOdd == sumEven) {
        printf("Yes, sum = %lf", sumOdd);
    } else {
        printf("No, diff = %lf", fabs(sumOdd - sumEven));
    }
}

void VowelsSum() {
    char input[256];
    int sum = 0;

    fgets(input, sizeof(input), stdin);

    for (int i = 0; i < strlen(input); ++i) {

        switch (input[i]) {
            case 'a':
                sum += 1;
                break;
            case 'e':
                sum += 2;
                break;
            case 'i':
                sum += 3;
                break;
            case 'o':
                sum += 4;
                break;
            case 'u':
                sum += 5;
                break;
        }
    }

    printf("%d", sum);
}

void CleverLily() {
    int age, toyPrice;
    double washingMachinePrice;
    double money = 0, moneyFromToys = 0, moneyFromBirthdaysSum = 0;
    double totalMoney = 0;

    scanf("%d", &age);
    scanf("%lf", &washingMachinePrice);
    scanf("%d", &toyPrice);

    for (int i = 1; i <= age; ++i) {

        if (i % 2 == 0) {
            moneyFromBirthdaysSum += 10;
            money += moneyFromBirthdaysSum - 1;
        } else {
            moneyFromToys += toyPrice;
        }
    }

    totalMoney = money + moneyFromToys;

    if (totalMoney >= washingMachinePrice) {
        printf("Yes! %.2lf", totalMoney - washingMachinePrice);
    } else {
        printf("No! %.2lf", washingMachinePrice - totalMoney);
    }
}

void DivideWithoutReminder() {
    int initialNumber, number, p1 = 0, p2 = 0, p3 = 0;
    scanf("%d", &initialNumber);

    for (int i = 1; i <= initialNumber; ++i) {
        scanf("%d", &number);

        if (number % 2 == 0) {
            p1++;
        }

        if (number % 3 == 0) {
            p2++;
        }

        if (number % 4 == 0) {
            p3++;
        }
    }

    double p1Percentage = (double) p1 / initialNumber * 100;
    double p2Percentage = (double) p2 / initialNumber * 100;
    double p3Percentage = (double) p3 / initialNumber * 100;
    
    printf("%.2lf%%\n", p1Percentage);
    printf("%.2lf%%\n", p2Percentage);
    printf("%.2lf%%\n", p3Percentage);
}