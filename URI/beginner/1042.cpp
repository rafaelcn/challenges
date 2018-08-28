#pragma GCC optimize("Ofast")

#include <iostream>
#include <string>
#include <cstdio>
#include <algorithm>
#include <vector>

int main() {

    std::vector<int> a(3);

    scanf("%d %d %d", &a[0], &a[1], &a[2]);
    std::vector<int> sorted = a;
    std::sort(sorted.begin(), sorted.end());

    for (auto e : sorted) {
        std::cout << e << "\n";
    }

    std::cout << "\n";

    std::cout << a[0] << "\n" << a[1] << "\n" << a[2] << "\n";


    return 0;
}
