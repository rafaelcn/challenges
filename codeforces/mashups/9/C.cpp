#include <bits/stdc++.h>

using namespace std;

// C. Treino do BamBam

int main() {

	int n = 0;
    pair<int, int> max = {0, 0};

	cin >> n;

    vector<int> a(3);

	for (int i = 0; i < n; i++) {
		int ai;

		cin >> ai;

		a[i%3] += ai;

		if (a[i%3] > max.first) {
            max.first = a[i%3];
			max.second = i%3;
		}
	}

    //cout << max << a[0] << a[1] << a[2] << endl;

	switch (max.second) {
	case 0:
		cout << "chest\n";
        break;
	case 1:
		cout << "biceps\n";
        break;
	case 2:
		cout << "back\n";
        break;
	}

    return 0;
}
