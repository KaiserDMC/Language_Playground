#include <stdio.h>

void printTemperature();

void printTempCelsius();

void printTemperatureReverse();

void printEOF();

void countFromStdIn();
void replaceMultipleSpaces(char *str);

int main(void) {
    printTemperature();
    printTempCelsius();
    printTemperatureReverse();
    printEOF();
//    countFromStdIn();
    replaceMultipleSpaces("test string  one");

    return 0;
}

#define LOWER 0
#define UPPER 300
#define STEP 20

void printTemperature() {
    int fahr;

    printf("Fahrenheit to Celsius conversion table\n");
    for (fahr = LOWER; fahr <= UPPER; fahr = fahr + STEP) {
        printf("%3d %6.1f\n", fahr, (5.0 / 9.0) * (fahr - 32));
    }
}

void printTempCelsius() {
    int celsius;

    printf("Celsius to Fahrenheit conversion table\n");
    for (celsius = LOWER; celsius <= UPPER; celsius = celsius + STEP) {
        printf("%3d %6.1f\n", celsius, (9.0 / 5.0) * celsius + 32.0);
    }
}

void printTemperatureReverse() {
    int fahr;

    printf("Fahrenheit to Celsius conversion table - Reversed\n");
    for (fahr = UPPER; fahr >= LOWER; fahr = fahr - STEP) {
        printf("%3d %6.1f\n", fahr, (5.0 / 9.0) * (fahr - 32));
    }
}

void printEOF() {
    printf("%d\n", EOF);
}

void countFromStdIn() {
    int c, sp, tab, nl;
    sp = nl = tab = 0;

    while ((c = getchar()) != EOF) {
        if (c == ' ')
            ++sp;

        if (c == '\t')
            ++tab;

        if (c == '\n')
            ++nl;
    }

    printf("Space -> %d\nTab -> %d\nNewLine -> %d\n", sp, tab, nl);
}

void replaceMultipleSpaces(char *str) {
    char *dest = str;

    while (*str != '\0') {
        while (*str == ' ' && *(str + 1) == ' ')
            str++;

        *dest++ = *str++;
    }

    *dest = '\0';

    printf("%s", dest);
}