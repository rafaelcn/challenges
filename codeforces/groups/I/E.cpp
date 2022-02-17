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

signed main() {
    ios_base::sync_with_stdio(false);
    cin.tie();

    int n;
    in(n);

    vector<int> a(n);

    for (int i = 0; i < n; i++) {
        in(a[i]);
    }

    sort(all(a));

    int d, i, j;

    i = 0;
    j = 1;
    d = 0;

    do {
		if (j > a[i]) {
            a.erase(a.begin() + i);
		} else {
			a.erase(a.begin() + i);

			i = 0;
			j++;
			d++;
		}

        i++;
	} while (i < a.size());

    out(d);
}