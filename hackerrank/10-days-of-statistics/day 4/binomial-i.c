#include <stdio.h>
#include <math.h>
#include <stdlib.h>

unsigned long int factorial(int, unsigned long int);
unsigned long int binomial(int, int);

int main(int argc, char **argv) {

    // number of trials
    int n = 6;

    // probability of a boy
    float p = 1.09 / (1+1.09);
    // probability of a girl
    float q = 1 - p;
    // result
    float  r = 0.0;

    for (int i = 3; i <= 6; i++) {
        r += binomial(n, i)*pow(p, i)*pow(q, (n-i));
    }

    printf("%.3f\n", r);

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

