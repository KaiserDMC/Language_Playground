#include <stdio.h>

void printHello();

void printTemperature();

void printTemperatureVersion2();

void printTemperatureVersionFor();

void printTemperatureVersionConsts();

void copyFromStdIn();

void copyFromStdInVersion2();

void charCountFromStdIn();

void charCountFromStdInVersion2();

void rowsCountFromStdIn();

void simpleWordCount();

void countersWithArrays();

int power(int m, int n);

void longestInputLine();

main() {
    printHello();
    printTemperature();
    printTemperatureVersion2();
    printTemperatureVersionFor();
    printTemperatureVersionConsts();
    copyFromStdIn();
    copyFromStdInVersion2();
    charCountFromStdIn();
    charCountFromStdInVersion2();
    rowsCountFromStdIn();
    simpleWordCount();
    countersWithArrays();

    printf("%d %d %d\n", 10, power(2, 10), power(-3, 10));
    
    longestInputLine();

    return 0;
}

void printHello() {
    printf("hello, world\n");
}

void printTemperature() {
    int fahr, celsius;
    int lower, upper, step;

    lower = 0;
    upper = 300;
    step = 20;

    fahr = lower;

    while (fahr <= upper) {
        celsius = 5 * (fahr - 32) / 9;
        printf("%d\t%d\n", fahr, celsius);
        fahr = fahr + step;
    }
}

void printTemperatureVersion2() {
    float fahr, celsius;
    float lower, upper, step;

    lower = 0;
    upper = 300;
    step = 20;

    fahr = lower;
    while (fahr <= upper) {
        celsius = (5.0 / 9.0) * (fahr - 32.0);
        printf("%3.0f %6.1f\n", fahr, celsius);
        fahr = fahr + step;
    }
}

void printTemperatureVersionFor() {
    int fahr;

    for (fahr = 0; fahr <= 300; fahr = fahr + 20) {
        printf("%3d %6.1f\n", fahr, (5.0 / 9.0) * (fahr - 32));
    }
}

#define LOWER 0
#define UPPER 300
#define STEP 20

void printTemperatureVersionConsts() {
    int fahr;

    for (fahr = LOWER; fahr <= UPPER; fahr = fahr + STEP) {
        printf("%3d %6.1f\n", fahr, (5.0 / 9.0) * (fahr - 32));
    }
}

void copyFromStdIn() {
    int c;

    c = getchar();
    while (c != EOF) {
        putchar(c);
        c = getchar();
    }
}

void copyFromStdInVersion2() {
    int c;

    while ((c = getchar()) != EOF) {
        putchar(c);
    }
}

void charCountFromStdIn() {
    long nc;
    nc = 0;

    while (getchar() != EOF)
        ++nc;
    printf("%ld\n", nc);
}

void charCountFromStdInVersion2() {
    double nc;

    for (nc = 0; getchar() != EOF; ++nc);

    printf("%.0f\n", nc);
}

void rowsCountFromStdIn() {
    int c, nl;
    nl = 0;

    while ((c = getchar()) != EOF)
        if (c == '\n')
            ++nl;
    printf("%d\n", nl);
}

void wordsCountFromStdIn() {
    int c, nl;
    nl = 0;

    while ((c = getchar()) != EOF)
        if (c == '\n')
            ++nl;
    printf("%d\n", nl);
}

#define IN 1
#define OUT 0

void simpleWordCount() {
    int c, nl, nw, nc, state;

    state = OUT;
    nl = nw = nc = 0;

    while ((c = getchar()) != EOF) {
        ++nc;
        if (c == '\n')
            ++nl;
        if (c == ' ' || c == '\n' || c == '\t')
            state = OUT;
        else if (state == OUT) {
            state = IN;
            ++nw;
        }
    }

    printf("%d %d %d\n", nl, nw, nc);
}

void countersWithArrays() {
    int c, i, nwhite, nother;
    int ndigit[10];

    nwhite = nother = 0;
    for (i = 0; i < 10; ++i)
        ndigit[i] = 0;

    while ((c = getchar()) != EOF)
        if (c >= '0' && c <= '9')
            ++ndigit[c - '0'];
        else if (c == ' ' || c == '\n' || c == '\t')
            ++nwhite;
        else
            ++nother;

    printf("digits =");
    for (i = 0; i < 10; ++i)
        printf(" %d", ndigit[i]);
    printf(", white space = %d, other = %d\n", nwhite, nother);
}

int power(int base, int n) {
    int i, p;

    p = 1;
    for (i = 1; i <= n; ++i)
        p = p * base;
    return p;
}

#define MAXLINE 1000

void longestInputLine() {
    int cgetline(char s[], int lim);
    void copy(char to[], char from[]);

    int len, max;
    char line[MAXLINE];
    char longest[MAXLINE];
    max = 0;

    while ((len = cgetline(line, MAXLINE)) > 0)
        if (len > max) {
            max = len;
            copy(longest, line);
        }
    if (max > 0)
        printf("%s", longest);
}

int cgetline(char s[], int lim) {
    int c, i;

    for (i = 0; i < lim - 1 && (c = getchar()) != EOF && c != '\n'; ++i)
        s[i] = c;
    if (c == '\n') {
        s[i] = c;
        ++i;
    }
    s[i] = '\0';
    return i;
}

void copy(char to[], char from[]) {
    int i;
    i = 0;

    while ((to[i] = from[i]) != '\0')
        ++i;
}