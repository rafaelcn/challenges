#include <bits/stdc++.h>

using namespace std;

using ll = long long;

#define all(x) begin(x), end(x)

int32_t main() {
    ios::sync_with_stdio(false); cin.tie(0); cout.tie(0);

    int n, error;
    multiset<int> errors_0, errors_1, errors_2, diff;

    cin >> n;

    for (int i = 0; i < n; i++) {
        cin >> error;
        errors_0.insert(error);
    }

    // second line
    for (int i = 0; i < n-1; i++) {
        cin >> error;
        errors_1.insert(error);
    }

    // third line
    for (int i = 0; i < n-2; i++) {
        cin >> error;
        errors_2.insert(error);
    }

    // O(n)
    //set_difference(all(errors), all(errors_1), inserter(diff, diff.begin()));
    set_difference(all(errors_0), all(errors_2), inserter(diff, diff.begin()));

    for (auto e : diff) {
        cout << e << "\n";
    }

}