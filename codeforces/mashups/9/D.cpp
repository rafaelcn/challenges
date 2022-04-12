#include <bits/stdc++.h>

using namespace std;

// D. Jogo de cartas pra dois

int main() {

    typedef long long ll;

	int n;
	cin >> n;

	vector<ll> a(n), b(2);

	for (int i = 0; i < n; i++) {
		cin >> a[i];
	}

	sort(a.begin(), a.end(), greater<ll>());

	for (int i = 0; i < n; i++) {
		b[i%2] += a[i];
	}

    //cout << a << "\n" << b << endl;

	cout << (b[0]-b[1]) << "\n";

    return 0;
}
