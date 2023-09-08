#include <stdio.h>
#include <limits.h>
#include <string.h>
#include <stdlib.h>
#include <math.h>

void NumberInRange();

void Sequence2Kp1();

void AccountBalance();

void MaxNumber();

void MinNumber();

void Graduation();

void GraduationPt2();

void Moving();

void Cake();

int main() {
    Cake();
    return 0;
}

void NumberInRange() {
    int number;

    while (1) {
        scanf("%d", &number);

        if (number >= 1 && number <= 100) {
            printf("The number is: %d\n", number);
            break;
        } else {
            printf("Invalid number!\n");
        }
    }
}

void Sequence2Kp1() {
    int initialNumber, number = 1;

    scanf("%d", &initialNumber);

    while (1) {
        printf("%d\n", number);

        number = 2 * number + 1;

        if (number > initialNumber) {
            break;
        }
    }
}

void AccountBalance() {
    int initialNumber;
    double number, balance = 0;

    scanf("%d", &initialNumber);

    while (1) {
        scanf("%lf", &number);

        if (number < 0) {
            printf("Invalid operation!\n");
            break;
        }

        balance += number;

        printf("Increase: %.2lf\n", number);

        initialNumber--;
        if (initialNumber == 0) {
            break;
        }
    }

    printf("Total: %.2lf\n", balance);
}

void MaxNumber() {
    int initialNumber, number, maxNumber = INT_MIN;

    scanf("%d", &initialNumber);

    while (1) {
        scanf("%d", &number);

        if (number > maxNumber) {
            maxNumber = number;
        }

        initialNumber--;
        if (initialNumber == 0) {
            break;
        }
    }

    printf("%d\n", maxNumber);
}

void MinNumber() {
    int initialNumber, number, minNumber = INT_MAX;

    scanf("%d", &initialNumber);

    while (1) {
        scanf("%d", &number);

        if (number < minNumber) {
            minNumber = number;
        }

        initialNumber--;
        if (initialNumber == 0) {
            break;
        }
    }

    printf("%d\n", minNumber);
}

void Graduation() {
    char name[100];
    int counter = 0;
    double grade, sum = 0;

    scanf("%s", name);

    while (1) {
        scanf("%lf", &grade);

        if (grade >= 4) {
            sum += grade;
            counter++;
        }

        if (counter == 12) {
            break;
        }
    }

    printf("%s graduated. Average grade: %.2lf\n", name, sum / 12);
}

void GraduationPt2() {
    char name[100];
    int counter = 0, failed = 0;
    double grade, sum = 0;

    scanf("%s", name);

    while (1) {
        scanf("%lf", &grade);

        if (grade >= 4) {
            sum += grade;
            counter++;
        } else {
            failed++;
        }

        if (failed > 1) {
            printf("%s has been excluded at %d grade\n", name, counter + 1);
            break;
        }

        if (counter == 12) {
            printf("%s graduated. Average grade: %.2lf\n", name, sum / 12);
            break;
        }
    }
}

void Moving() {
    int width, length, height, volume = 0, freeSpace = 0;
    scanf("%d %d %d", &width, &length, &height);

    volume = width * length * height;

    while (1) {
        char input[100];
        scanf("%s", input);

        if (strcmp(input, "Done") == 0) {
            break;
        }

        int boxes = atoi(input);
        freeSpace += boxes;

        if (freeSpace > volume) {
            printf("No more free space! You need %d Cubic meters more.\n", freeSpace - volume);
            break;
        }
    }

    if (freeSpace <= volume) {
        printf("%d Cubic meters left.\n", volume - freeSpace);
    }
}

void Cake() {
    int width, length, cakeSize = 0;

    scanf("%d %d", &width, &length);
    cakeSize = width * length;

    while (1) {
        char input[100];
        scanf("%s", input);

        if (strcmp(input, "STOP") == 0) {
            if (cakeSize > 0) {
                printf("%d pieces are left.", cakeSize);
            }
            break;
        }

        int pieces = atoi(input);

        cakeSize -= pieces;
        
        if (cakeSize < 0) {
            printf("No more cake left! You need %g pieces more.", fabs(cakeSize));
            break;
        }
    }
}