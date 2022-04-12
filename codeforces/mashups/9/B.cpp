#include <bits/stdc++.h>

using namespace std;

// B. Números menores ou iguais

int main() {
    ios_base::sync_with_stdio(false); cin.tie(0); cout.tie(0);

    int n, m;

    cin >> n >> m;

    vector<int> a(n), b(m);

    for (int i = 0; i < n; i++) {
        cin >> a[i];
    }

    for (int i = 0; i < m; i++) {
        cin >> b[i];
    }

    sort(a.begin(), a.end());

    for (int i = 0; i < m; i++) {
        auto it = upper_bound(a.begin(), a.end(), b[i]);
        cout << (it - a.begin()) << " ";
    }

    cout << "\n";

    return 0;
}