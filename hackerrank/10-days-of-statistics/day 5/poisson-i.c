#include <stdio.h>
#include <math.h>
#include <stdlib.h>

int fact(int a, int b);

int main() {

    float mean = 2.5;
    float p = 5;

    float rp = pow(mean, p)/(fact(5, 1)*exp(mean));

    printf("%.3f", rp);

    return 0;
}

int fact(int a, int b) {
    if (a == 0) {
        return b;
    }

    return fact(a-1, b*a);
}