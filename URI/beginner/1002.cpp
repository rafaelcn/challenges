#include <iostream>
#include <cstdio>

#define PI 3.14159

int main() {
	
	using namespace std;

	double area, r;

	scanf("%lf", &r);

	area = PI*r*r;

	printf("A=%.4lf", area);

	return 0;
}
