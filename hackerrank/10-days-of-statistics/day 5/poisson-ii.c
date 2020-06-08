#include <stdio.h>
#include <math.h>
#include <stdlib.h>

int main(int argc, char **argv) {

    float m = 0.88;
    float m1 = 1.55;

    float expected_1 = 160 + 40 * (m + pow(m, 2));
    float expected_2 = 128 + 40 * (m1 + pow(m1, 2));

    printf("%.3f\n", expected_1);
    printf("%.3f\n", expected_2);

    return 0;
}