#include <bits/stdc++.h>
#include <iomanip>

int main() {
	using namespace std;

	int hours, id;
	float per_hour;

	cin >> id >> hours >> per_hour;

	cout << "NUMBER = " << id << endl;
	cout << fixed << setprecision(2);
	cout << "SALARY = U$ " << (hours*per_hour) << endl;

	return 0;
}
