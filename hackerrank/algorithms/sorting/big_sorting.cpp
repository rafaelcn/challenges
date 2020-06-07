#pragma GCC optimize ("Ofast")

#include <bits/stdc++.h>

int main() {
    using namespace std;

    int n;
    cin >> n;

    vector<string> to_sort(n);

    for (int i = 0; i < n; i++) {
        cin >> to_sort[i];
    }

    // Sort using the lexicographical order of the string.
    // Converting from/to a unsigned long long number isn't feasible due
    // to the problem constraint.
    sort(to_sort.begin(), to_sort.end(), [](const string a, const string b) {
            int sa = a.size();
            int sb = b.size();

            if (sa == sb) {
                return sb < sa;
            }

            return b < a;
    });

    for (const auto& e: to_sort) {
        cout << e << "\n";
    }

    return 0;
}
