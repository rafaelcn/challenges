#pragma GCC optimize("Ofast")

#include <iostream>
#include <string>
#include <cstdio>
#include <algorithm>

long long recursive = 0;

long long fib(int);

int main() {

    using namespace std;

    int n, d;
    cin >> d;

    while (d > 0) {
        cin >> n;
        cout << "fib(" << n << ") = " << recursive-1 << " calls = " << fib(n) << "\n";
        recursive = 0;
        d--;
    }

    return 0;
}

long long fib(int n) {

    recursive+=1;

    if (n == 0) {
        return 0;
    } else if (n == 1) {
        return 1;
    }

    return fib(n-1) + fib(n-2);
}
