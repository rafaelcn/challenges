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


int32_t main() {

    ll n, q; in(n, q);

    vector<ll> a(1e5+2, 0);
    vector<ll> d(1e5+2, 0);

    while (q > 0) {
        ll l, r, x;
        in(l, r, x);

        l -= 1;
        r -= 1;

        d[l] += x;
        d[r+1] -= x;

        q--;
    }

    ll v = 0;
    for (int i = 0; i <= n; i++) {
        v += d[i];
        a[i] = v;
    }

    for (int i = 0; i < n; i++) {
        cout << a[i] << " ";
    }
    out("\n");

    return 0;
}