#include <iostream>

using namespace std;

int main() {

    int n;

    cin >> n;

    while (n > 0) {
        long int n1;
        bool negative = false;

        cin >> n1;

        if (n1 == 0) {
            cout << "NULL\n";
        } else {
            if (n1 < 0) {
                negative = true;
            }

            if ((n1 & 0x1) == 1) {
                if (negative) {
                    cout << "ODD NEGATIVE\n";
                } else {
                    cout << "ODD POSITIVE\n";
                }
            } else {
                if (negative) {
                    cout << "EVEN NEGATIVE\n";
                } else {
                    cout << "EVEN POSITIVE\n";
                }
            }
        }

        --n;
    }

    return 0;
}