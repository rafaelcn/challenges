#include <bits/stdc++.h>

using namespace std;

int32_t main() {
    ios::sync_with_stdio(false);
    cin.tie();
    cout.tie();

    int q, qt, qn;
    set<int> numbers;

    cin >> q;

    while (q > 0) {
        cin >> qt >> qn;

        if (qt == 1) {
            numbers.insert(qn);
        } else {
            auto it = numbers.find(qn);
            if (it != numbers.end()) {
                cout << "Sim\n";
            } else {
                cout << "Nao\n";
            }
        }

        q--;
    }

    return 0;
}