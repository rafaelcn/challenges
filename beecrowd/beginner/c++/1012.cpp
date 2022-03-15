#include <bits/stdc++.h>

using namespace std;

int32_t main() {
    double a, b, c; cin >> a >> b >> c;

    const double pi = 3.14159;

    printf("TRIANGULO: %.3f\n", a*c/2);
    printf("CIRCULO: %.3f\n", c*c*pi);
    printf("TRAPEZIO: %.3f\n", (a+b)*c/2);
    printf("QUADRADO: %.3f\n", b*b);
    printf("RETANGULO: %.3f\n", a*b);

    return 0;
}