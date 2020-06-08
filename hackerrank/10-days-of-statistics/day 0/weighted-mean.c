#include <stdio.h>
#include <math.h>
#include <stdlib.h>

int sum(float v[], int n) {
    float s = 0.0;

    for (int i = 0; i < n; i++) {
        s += v[i];
    }

    return s;
}

int main() {

    int n;

    scanf("%d", &n);

    float x[n];
    float w[n];
    float wmean = 0.0;

    for (int i = 0; i < n; i++) {
        scanf("%f", &x[i]);
    }

    for (int i = 0; i < n; i++) {
        scanf("%f", &w[i]);
    }

    float sw = sum(w, n);

    for (int i = 0; i < n; i++) {
        wmean += (x[i]*w[i])/sw;
    }

    printf("%.1f\n", wmean);

    return 0;
}
