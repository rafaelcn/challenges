#pragma GCC optimize("Ofast")

#include <bits/stdc++.h>

using namespace std;

long int solve(long int n) {
    long int z = 0;

    // n != 0
    while (n) {
        if ((n & 1) == 0) {
            z++;
        }
        n >>= 1;
    }

    return z;
}

int main() {
    long int n;
    scanf("%ld", &n);

    long int r = solve(n);

    printf("%ld", r);

}
