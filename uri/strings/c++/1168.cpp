#include <iostream>
#include <string>
#include <cstdio>
#include <cstdlib>
#include <algorithm>

int lq(int);

int main() {

	char *s = (char *) malloc(sizeof(10E100));
	int n = 0;

	scanf("%d", &n);


	while (n > 0) {
		scanf("%s", s);
		long result = 0;

		while (*s) {
			result += lq(*s-'0');
			s++;
		}

		std::cout << result << " leds" << "\n";
		n--;
	}

	return 0;
}

int lq(int led) {
	int v[10] = {6, 2, 5, 5, 4, 5, 6, 3, 7, 6};

	return v[led];
}