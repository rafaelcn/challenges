#include <bits/stdc++.h>

using namespace std;

#pragma GCC optimize("Ofast")
#pragma GCC target("avx,avx2,fma")

using ii = pair<int, int>;
using ll = long long;
using dd = double;

#define sz(v) (int)v.size()
#define all(x) x.begin(), x.end()
#define apply(v, f) for_each(all(v), f)
#define cond(c, t, f) ((c) ? (t) : (f))
#define loop(ii, n) for (int ii = 0; ii < (n); ++ii)

// Container IO
template <typename T> istream &operator>>(istream &is, vector<T> &v) { for (auto &a : v) is >> a; return is; }
template <typename T> ostream &operator<<(ostream &os, vector<T> v) {
  loop(i, sz(v)) os << cond(i," ","") << v[i]; return os; }
template <typename T> ostream &operator<<(ostream &os, set<T> v) {
  for (auto it = v.begin(); it != v.end(); it++) os << cond(it != v.begin()," ","") << *it; return os; }
template <typename F, typename S> ostream &operator<<(ostream &os, pair<F, S> v) {
  os << v.first << " " << v.second; return os; }

// General IO
template <typename... A> void in(A &...a) { ((cin >> a), ...); }
template <typename... A> void out(A... a) { ((cout << a << " "), ...); cout << endl; }
template <typename... A> void print(A... a) { ((cout << a), ...); }

void solve() {

    ll n;
	in(n);

	if (n == 1) {
		out("1");
	} else if (n == 2) {
		out("2");
	} else {
		// combination of ones an twos
		// naive solution (start both strings an choose one)
		string out1, out2;
		ll sum1, sum2;

		bool toggle = true;

		while (sum1 <= n || sum2 <= n) {
			if (toggle) {
				sum1 += 1;
				sum2 += 2;
				out1 += "1";
				out2 += "2";
			} else {
				sum1 += 2;
				sum2 += 1;
				out1 += "2";
				out2 += "1";
			}

			if (sum1 == n || sum2 == n) {
				break;
			}

			toggle = !toggle;
		}

		if (sum1 == sum2) {
			out(sum2);
		} else {
			if (sum1 == n) {
				out(out1, sum1);
			} else if (sum2 == n) {
				out(out2, sum2);
			}
		}
	}
}

int32_t main() {

    ll t; in(t);

    while (t > 0) {
        solve();
        t--;
    }

    return 0;
}