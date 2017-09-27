#pragma GCC optimize("Ofast")

#include <bits/stdc++.h>

using namespace std;

bool ab(vector<int>& v, int a, int b);
bool ba(vector<int>& v, int a, int b);

int main() {
	
	int n, q;
	cin >> n;
	
	vector<int> s(n);

	for (int i = 0; i < n; i++) {
		cin >> s[i];
	}
	
	
	char c;
	int a = 0;
	int b = 0; 

	cin >> q;

	while (q > 0) {
		cin >> c >> a >> b;

		if (c == 'Q') {
			if (ab(s, a, b)) {
				printf("A\n");
			}
			else if (ba(s, a, b)) {
				printf("D\n");
			} else {
				printf("X\n");
			}
		} else {
			int t = s[a-1];
			s[a-1] = s[b-1];
			s[b-1] = t;
		} 

		q--;
	}
	

    return 0;
}

bool ab(vector<int>& v, int a, int b) {
	for (int i = a-1; i < b-1; i++) {
		if (v[i] >= v[i+1]) {
			return false;
		}
	}

	return true;
}

bool ba(vector<int>& v, int a, int b) {
	for (int i = a-1; i < b-1; i++) {
		if (v[i] <= v[i+1]) {
			return false;
		}
	}

	return true;
}


