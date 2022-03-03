#include <bits/stdc++.h>

using namespace std;

using ll = long long;

#define all(x) begin(x), end(x)

int32_t main() {
    ios::sync_with_stdio(false); cin.tie(0); cout.tie(0);

    int n, error;

    cin >> n;

    vector<ll> a(n), b(n-1), c(n-2);

    for (int i = 0; i < n; i++) {
        cin >> error;
        a[i] = error;
    }

    // second line (which I don't care as the fixed error will be included later)
    for (int i = 0; i < n-1; i++) {
        cin >> error;
    }

    // third line
    for (int i = 0; i < n-2; i++) {
        cin >> error;
        c[i] = error;
    }

    sort(all(a));
    sort(all(c));

    int j = 0;
    vector<int> remaining;

    // 1 3 3 4 5 7
    //   3 3 4 5 7 -> fixed one error
    //     3 4 5 7 -> fixed two errors

    for (int i = 0; i < n; i++) {
        if (a[i] != c[j]) {
            remaining.push_back(a[i]);
        } else {
           j++;
        }
    }

    for (auto e : remaining) {
        cout << e << "\n";
    }
}