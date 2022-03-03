#include <bits/stdc++.h>

using namespace std;

using ll = long long;

#define all(x) begin(x), end(x)

int32_t main() {
    ios::sync_with_stdio(false); cin.tie(0); cout.tie(0);

    int n, k1, k2, bi;

    cin >> n >> k1 >> k2;

    vector<ll> a(n), b(n);
    priority_queue<ll> c(n, 0);

    for (int i = 0; i < n; i++) {
        cin >> a[i];
    }

    for (int i = 0; i < n; i++) {
        cin >> b[i];
        c.push(abs(a[i]-b[i]));
    }

    int k = k1+k2;

    while (k > 0) {
        ll t = c.top();

        c.pop();
        c.push(abs(t-1));

        k--;
    }

    ll cost = 0;
    while (!c.empty()) {
        ll ci = c.top();
        c.pop();

        cost += ci*ci;
    }

    cout << cost << "\n";
}