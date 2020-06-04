#include <bits/stdc++.h>
#include <iomanip>

int main() {
	using namespace std;

	string name;
	double salary, sold;

	cin >> name >> salary >> sold;

	cout << fixed << setprecision(2);
	cout << "TOTAL = R$ " << salary+(sold*0.15) << endl;

	return 0;
}
