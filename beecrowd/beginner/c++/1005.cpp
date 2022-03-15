#include <iostream>
#include <iomanip>

int main() {
	float a, b;
	double average;
	std::cin >> a >> b;

	a = a*3.5;
	b = b*7.5;

	average = (a+b)/11;
	std::cout << std::fixed;
	std::cout << "MEDIA = " << std::setprecision(5) << average << std::endl;
	return 0;
}
