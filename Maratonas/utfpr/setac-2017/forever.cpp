#pragma GCC optimize("Ofast")

#include <bits/stdc++.h>

using namespace std;

int main() {
	int n, p;
	
	cin >> n >> p;
	
	bool forever = true;
	
	// problema de fatoração
	for (int i = 2; i < n; i++) {
		if (((p % i == 0) or (i % p == 0)) and p != i) {
			forever = false;
			break;
		}
	} 

	if (forever) {
		printf("sim\n");
	} else {
		printf("nao\n");
	}

    return 0;
}
