#include <iostream>
#include <cmath>

int main() {
    int a;
    int b;
    int sum = 0;

    std::cin >> a >> b;

    for (int i = std::min(a, b); i <= std::max(a, b); i++) {
        if (i % 13 != 0) {
            sum += i;
        }
    }

    std::cout << sum << std::endl;

    return 0;
}
