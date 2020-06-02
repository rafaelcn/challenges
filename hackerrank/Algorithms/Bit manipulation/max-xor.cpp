#include <bits/stdc++.h>

using namespace std;

int max_xor(int l, int r) {
    int n = l^r;
    int s = 0;

    while (n) {
        s++;
        n >>= 1;
    }

    return pow(2, s)-1;
}

int main() {
    int l, r;

    cin >> l >> r;

    int res = max_xor(l, r);
    printf("%d", res);

    return 0;
}
