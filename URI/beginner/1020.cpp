#include <iostream>
#include <cstdio>

int main() {

    int n;

    scanf("%d", &n);

    int ry = n % 365;
    int y = (n-ry) / 365;

    int rm = (ry) % 30;
    int m = (ry-rm) / 30;

    int d = n-y*365-m*30;

    printf("%d ano(s)\n%d mes(es)\n%d dia(s)\n", y, m, d);


    return 0;
}
