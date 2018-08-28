#pragma GCC optimize("Ofast")

#include <iostream>
#include <string>
#include <cstdio>
#include <algorithm>

int main() {

    char *expr = (char *) malloc(sizeof(1000));

    while (scanf("%s", expr) != EOF) {
        int l = 0;

        char *s = expr;

        while (*s) {

            if (*s == '(') l++;
            if (*s == ')') l--;

            if (l < 0)  break;

            s++;
        }

        if (l != 0) {
            std::cout << "incorrect\n";
        } else {
            std::cout << "correct\n";
        }
    }

    return 0;
}
