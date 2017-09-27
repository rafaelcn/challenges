#pragma GCC optimize("Ofast")

#include <bits/stdc++.h>

using namespace std;

bool limpo(vector<int>& v);
bool limpo_posicao(vector<int>& v, int& i);

void inc(int& i, int& n);

int main() {
	int n, d, c = 0;
	int i = 1;
	int max = 0;	
	
	cin >> n >> d;

	vector<int> s(n);

	for (int i = 0; i < s.size(); i++) {
		cin >> s[i];
		if (s[i] >= max) {
			max = s[i];
		}
	}

	while(1) {
		if (limpo_posicao(s, i)) {
			inc(i, n);
		} else {

			s[i] - d < 0 ? (s[i] = 0) : (s[i] -= d);
			s[i-1] - d < 0 ? (s[i-1] = 0) : (s[i-1] -= d);
			s[i+1] - d < 0 ? (s[i+1] = 0) : (s[i+1] -= d);		

			inc(i, n);

			c++;
		}

		if (c > ceil(double(max/d))) { 
			if (limpo(s)) {	
				break;
			}
		}
	}

	printf("%d\n", c);

    return 0;
}

bool limpo(vector<int>& v) {
	bool limpo = true;

	// O(n)
	for (int i = 0; i < v.size(); i++) {
		if (v[i] != 0) {
			limpo = false;
			break;
		}
	}

	return limpo;
}

bool limpo_posicao(vector<int>& v, int& i) {
	bool limpo = true;

	if (v[i] != 0 or v[i-1] != 0 or v[i+1] != 0)  {
		limpo = false;
	}

	return limpo;
}

void inc(int& i, int &n)  {
	if (n % 3 == 0) {
		i += 3;
	} else {
		i += 2;
	}

	if (i >= n) {
		i = 1;
	}	
}
