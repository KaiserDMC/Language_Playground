#include <stdio.h>
#include <math.h>
#include <string.h>
#include <stdbool.h>

void NumbersNto1();

void Numbers1toN();

void EvenPowersOf2();

void Combinations();

void Building();

void Travelling();

void NameWars();

void CookieFactory();

void MagicNumbers();

void PasswordGenerator();

int main() {
    Travelling();
    return 0;
}

void NumbersNto1() {
    int n;
    scanf("%d", &n);

    for (int i = n; i >= 1; i--) {
        printf("%d\n", i);
    }
}

void Numbers1toN() {
    int n;
    scanf("%d", &n);

    for (int i = 1; i <= n; i += 3) {
        printf("%d\n", i);
    }
}

void EvenPowersOf2() {
    int n;
    scanf("%d", &n);

    for (int i = 0; i <= n; i += 2) {
        printf("%d\n", (int) pow(2, i));
    }
}

void Combinations() {
    int n;
    scanf("%d", &n);

    int counter = 0;

    for (int x1 = 0; x1 <= n; x1++) {
        for (int x2 = 0; x2 <= n; x2++) {
            for (int x3 = 0; x3 <= n; x3++) {
                for (int x4 = 0; x4 <= n; x4++) {
                    for (int x5 = 0; x5 <= n; x5++) {
                        if (x1 + x2 + x3 + x4 + x5 == n) {
                            counter++;
                        }
                    }
                }
            }
        }
    }

    printf("%d", counter);
}

void Building() {
    int floors;
    scanf("%d", &floors);

    int rooms;
    scanf("%d", &rooms);

    for (int i = floors; i >= 1; i--) {
        for (int j = 0; j < rooms; j++) {
            if (i == floors) {
                printf("L%d%d ", i, j);
            } else if (i % 2 == 0) {
                printf("O%d%d ", i, j);
            } else {
                printf("A%d%d ", i, j);
            }
        }
        printf("\n");
    }
}

void Travelling() {
    char destination[100];

    while (1) {
        fgets(destination, sizeof(destination), stdin);

        if (strcmp(destination, "End\n") == 0) {
            break;
        }

        double minimumBudget = 0.0;
        scanf("%lf", &minimumBudget);

        // Consume the newline character after reading minimumBudget
        getchar();

        size_t len = strlen(destination);
        if (len > 0 && destination[len - 1] == '\n') {
            destination[len - 1] = '\0';
        }
        double leftToSave = 0.0;

        double sumSaved;
        scanf("%lf", &sumSaved);

        // Consume the newline character after reading sumSaved
        getchar();

        leftToSave += sumSaved;

        while (leftToSave < minimumBudget) {
            scanf("%lf", &sumSaved);

            // Consume the newline character after reading sumSaved
            getchar();

            leftToSave += sumSaved;
        }

        if (leftToSave >= minimumBudget) {
            printf("Going to %s!\n", destination);
        }
    }
}

void NameWars() {
    char name[100];
    scanf("%s", name);

    int maxSum = 0;
    char maxName[100];

    while (strcmp(name, "STOP") != 0) {
        int sum = 0;
        for (int i = 0; i < strlen(name); i++) {
            sum += name[i];
        }

        if (sum > maxSum) {
            maxSum = sum;
            strcpy(maxName, name);
        }

        scanf("%s", name);
    }

    printf("Winner is %s - %d!", maxName, maxSum);
}

void CookieFactory() {
    int batches, counter = 0;
    bool has_flour = false, has_eggs = false, has_sugar = false;
    scanf("%d", &batches);

    while (counter < batches) {
        char ingredient[100];
        scanf("%s", ingredient);

        while (strcmp(ingredient, "Bake!") != 0) {

            if (strcmp(ingredient, "flour") == 0) {
                has_flour = true;
            } else if (strcmp(ingredient, "eggs") == 0) {
                has_eggs = true;
            } else if (strcmp(ingredient, "sugar") == 0) {
                has_sugar = true;
            }

            scanf("%s", ingredient);
        }

        if (has_flour && has_eggs && has_sugar) {
            counter++;

            printf("Baking batch number %d...\n", counter);
            has_flour = false;
            has_eggs = false;
            has_sugar = false;
        } else {
            printf("The batter should contain flour, eggs and sugar!\n");
        }

    }
}

void MagicNumbers() {
    int magicNumber;
    scanf("%d", &magicNumber);

    for (int i = 1; i <= 9; i++) {
        for (int j = 0; j <= 9; j++) {
            for (int k = 0; k <= 9; k++) {
                for (int l = 0; l <= 9; l++) {
                    for (int m = 0; m <= 9; m++) {
                        for (int n = 0; n <= 9; n++) {
                            if (i * j * k * l * m * n == magicNumber) {
                                printf("%d%d%d%d%d%d ", i, j, k, l, m, n);
                            }
                        }
                    }
                }
            }
        }
    }
}

void PasswordGenerator() {
    int n, l;
    scanf("%d %d", &n, &l);

    for (int i = 1; i <= n; i++) {
        for (int j = 1; j <= n; j++) {
            for (int k = 1; k <= l; k++) {
                for (int m = 1; m <= l; m++) {
                    for (int o = 1; o <= n; o++) {
                        if (o > i && o > j) {
                            printf("%d%d%c%c%d ", i, j, (char) (k + 96), (char) (m + 96), o);
                        }
                    }
                }
            }
        }
    }
}