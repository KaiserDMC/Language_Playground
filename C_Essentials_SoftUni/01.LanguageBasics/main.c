#include <stdio.h>
#include <math.h>

void printHello();
void Name();
void Apples();
void RectangleArea();
void PythagorasProblem();

int main() {
    PythagorasProblem();
    return 0;
}

void printHello() {
    printf("Hello SoftUni!\n");
}

void Name() {
    printf("My name is Kris\n");
}

void Apples() {
    int treeOne = 150;
    int treeTwo = 142;
    int treeThree = 127;
    
    int totalApples = treeOne + treeTwo + treeThree;
    
    printf("Apples gathered = %d\n", totalApples);
}

void RectangleArea(){
    int width;
    scanf("%d", &width);
    
    int height;
    scanf("%d", &height);
    
    int area = width * height;
    
    printf("Area %d.\n", area);
}

void PythagorasProblem(){
    int a;
    scanf("%d", &a);
    
    int b;
    scanf("%d", &b);  
    
    int c = sqrt(a * a + b * b);

    printf("Hypotenuse is %d.\n", c);
}