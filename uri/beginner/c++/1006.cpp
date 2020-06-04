#include <bits/stdc++.h>
#include <iomanip>

int main() {
	using namespace std;

	float a, b, c, average;

	cin >> a >> b >> c;

	a = a*2;
	b = b*3;
	c = c*5;

	average = (a+b+c)/10;

	cout << fixed;
	cout << setprecision(1);
	cout << "MEDIA = " << average << endl;

	return 0;
}
