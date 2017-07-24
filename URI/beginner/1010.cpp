#include <bits/stdc++.h>
#include <iomanip>

int main() {

	using namespace std;

	int id, id_1, n_unit, n_unit_1;
	double price, price_1;

	cin >> id >> n_unit >> price;
	cin >> id_1 >> n_unit_1 >> price_1;

	cout << fixed << setprecision(2);
	cout << "VALOR A PAGAR: R$ " << (n_unit*price)+(n_unit_1*price_1) << endl;

	return 0;
}
