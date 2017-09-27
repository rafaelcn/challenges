#include <bits/stdc++.h>

using namespace std;

int main() {
	int x1, y1, r1;
	int x2, y2, r2;
	double d;

	cin >> x1 >> y1 >> r1;
	cin >> x2 >> y2 >> r2;

	d = sqrt(pow(x2-x1, 2)+pow(y2-y1, 2));
	
	if (d < r1+r2) {
		cout << "guerra" << "\n";
	} else {
		cout << "paz" << "\n";
	}

	return 0;
}
