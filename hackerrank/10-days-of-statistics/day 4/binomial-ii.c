#include <stdio.h>
#include <math.h>
#include <stdlib.h>

unsigned long int factorial(int, unsigned long int);
unsigned long int binomial(int, int);

int main(int argc, char **argv) {

    int n = 10;

    float p = 0.88;
    float q = 0.12;

    float r1 = 0.0;
    float r2 = 0.0;

    for (int i = 8; i <= 10; i++) {
        r1 += binomial(n, i)*pow(p, i)*pow(q, (n-i));
    }

    for (int i = 0; i <= 8; i++) {
        r2+= binomial(n, i)*pow(p, i)*pow(q, (n-i));
    }

    printf("%.3f\n", r1);
    printf("%.3f\n", r2);

    return 0;
}


unsigned long int factorial(int n, unsigned long int acc) {
    if (n == 0) {
        return acc;
    }

    return factorial(n-1, acc*n);
}

unsigned long int binomial(int n, int x) {
    return factorial(n, 1)/(factorial(x, 1)*factorial(n-x, 1));
}
