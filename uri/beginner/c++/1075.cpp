#include <iostream>

using namespace std;

int main() {

    int n;

    cin >> n;

    int i = 0;

    while (i <= 10000) {
        if (i % n == 2) {
            cout << i << "\n";
        }
        ++i;
    }

    return 0;
}