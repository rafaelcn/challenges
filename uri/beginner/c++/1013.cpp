#include <iostream>
#include <algorithm>

int main() {
    int n, n1, n2;

    std::cin >> n >> n1 >> n2;

    int max = std::max(n, std::max(n1, n2));

    std::cout << max << " eh o maior\n";

    return 0;
}