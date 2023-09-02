#include <stdio.h>
#include <math.h>

void SquareArea();

void InchesToCentimeters();

void CircleAreaAndPerimeter();

void PetShop();

void YardGreening();

void GreetingsByName();

void ProjectsCreation();

void FishTank();

void CharityCampaign();

int main() {
    CharityCampaign();
    return 0;
}

void SquareArea() {
    int side;
    scanf("%d", &side);

    printf("%d", side * side);
}

void InchesToCentimeters() {
    float inches;
    scanf("%f", &inches);

    printf("%f", inches * 2.54);
}

void CircleAreaAndPerimeter() {
    float radius;
    scanf("%f", &radius);

    printf("%.2f\n", M_PI * radius * radius);
    printf("%.2f\n", 2 * M_PI * radius);
}

void PetShop() {
    int zooDogsCount, ownDogsCount;

    scanf("%d", &zooDogsCount);
    scanf("%d", &ownDogsCount);
    
    double result = zooDogsCount * 2.5 + ownDogsCount * 4;
    printf("%.2f lv.", result);
}

void YardGreening(){
    double squareMeters;
    scanf("%lf", &squareMeters);

    double price = squareMeters * 7.61;
    double discount = price * 0.18;
    double finalPrice = price - discount;

    printf("The final price is: %.2f lv.\n", finalPrice);
    printf("The discount is: %.2f lv.", discount);
}

void GreetingsByName(){
    char str[100];
    scanf("%s", str);
    
    printf("Hello, %s!", str);
}

void ProjectsCreation(){
    char architectName[100];
    int buildingsCount;

    scanf("%s", architectName);
    scanf("%d", &buildingsCount);

    int hours = buildingsCount * 3;

    printf("The architect %s will need %d hours to complete %d project/s.", architectName, hours, buildingsCount);
}

void FishTank(){
    int length, width, height; // in cm
    double percent;
    
    scanf("%d", &length);
    scanf("%d", &width);
    scanf("%d", &height);
    scanf("%lf", &percent);

    double volume = length * width * height; // in cm3
    double liters = volume * 0.001; // in liters
    double percentDecimal = percent * 0.01;
    double result = liters * (1 - percentDecimal);

    printf("%.3f", result);
}

void CharityCampaign(){
    int days, workers, cakes, waffles, pancakes;
    
    scanf("%d", &days);
    scanf("%d", &workers);
    scanf("%d", &cakes);
    scanf("%d", &waffles);
    scanf("%d", &pancakes);

    double cakesPrice = cakes * 45;
    double wafflesPrice = waffles * 5.80;
    double pancakesPrice = pancakes * 3.20;

    double total = (cakesPrice + wafflesPrice + pancakesPrice) * workers * days;
    double result = total - (total / 8);

    printf("%.2f", result);
}