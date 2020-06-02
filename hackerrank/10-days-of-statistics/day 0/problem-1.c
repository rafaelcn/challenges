#include <stdio.h>
#include <math.h>
#include <stdlib.h>

float cmean(int *a, int n);
float cmedian(int *a, int n);
int cmode(int *a, int n);

int cmp(const void *a, const void *b) {
    float c = *(const float *)a;
    float d = *(const float *)b;

    if (c > d) {
        return 1;
    } else {
        return -1;
    }

    return 0;
}

int main() {
    float mean = 0.0;
    float median = 0.0;
    int mode = 0;

    unsigned int n = 0;

    scanf("%d", &n);

    int a[n];

    for(int i = 0; i < n; i++) {
        scanf("%d", &a[i]);
    }

    qsort(a, n, sizeof(int), cmp);

    mean = cmean(a, n);
    median = cmedian(a, n);
    mode = cmode(a, n);

    printf("%.1f\n", mean);
    printf("%.1f\n", median);
    printf("%d\n", mode);

    return 0;
}

float cmean(int *a, int n) {
    float d = 0.0;

    for (int i = 0; i < n; i++) {
        d += (float)a[i]/(float)n;
    }

    return d;
}

float cmedian(int *a, int n) {
    float m = 0.0;

    // if n is odd
    if (n & 1) {
        m = a[(int)floor(n/2)];
    } else {
        m = (float)(a[n/2]+a[n/2-1])/2;
    }

    return m;
}

int cmode(int *a, int n) {
    // Array a already sorted.
    // Creates a n-sized array.
    int *f = calloc(n, sizeof(int));
    // Mode is considered to be the first number of the a array.
    int m = a[0];

    int j = 0;
    int k = 0;

    for (int i = 0; i < n; i++) {
        if (a[i] == m) {
            f[j] += 1;
        } else {
            m = a[i];
            j++;

            f[j] += 1;
        }

        if (f[j] > k && f[j] > 1)
            k = i;
    }

    return a[k];
}
