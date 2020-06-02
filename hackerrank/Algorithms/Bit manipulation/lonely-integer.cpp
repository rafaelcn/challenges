#pragma GCC optimize("Ofast")

#include <bits/stdc++.h>

using namespace std;

int lonelyinteger(vector <int> a) {
    int r = 0;
    vector<int> is_in;

    if (a.size() == 1) {
        return a[0];
    }

    for (size_t i = 0; i < a.size(); i++) {

    }

    return r;
}

int main() {
    int n;
    cin >> n;

    vector<int> a(n);

    for(int i = 0; i < n; i++){
        cin >> a[i];
    }

    int result = lonelyinteger(a);
    printf("%d", result);

    return 0;
}
