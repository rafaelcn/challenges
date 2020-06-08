#include <stdio.h>
#include <math.h>
#include <stdlib.h>

int main(int argc, char **argv) {

    int n;
    float mean = 0.0;
    float variance = 0.0;

    scanf("%d", &n);

    int x[n];


    for (int i = 0; i < n; i++) {
        scanf("%d", &x[i]);
        mean += x[i]/n;
    }

    for (int i = 0; i < n; i++) {
        variance += pow(mean-x[i], 2);
    }

    float std = sqrt(variance/n);

    printf("%.1f\n", std);

    return 0;
}
