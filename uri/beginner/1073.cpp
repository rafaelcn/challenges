#pragma GCC optimize("Ofast")

#include <iostream>
#include <string>
#include <cstdio>
#include <algorithm>

int main() {

    using namespace std;

    long n = 0;

    cin >> n;

    for (int i = 2; i <= n; i+=2) {
        cout << i << "^" << 2 << " = " << i*i << "\n";
    }

    return 0;
}
