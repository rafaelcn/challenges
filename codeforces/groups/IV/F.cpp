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
template <typename T> ostream &operator<<(ostream &os, vector<T> v) { loop(i, sz(v)) os << cond(i," ","") << v[i]; cout << endl; return os; }
template <typename T> ostream &operator<<(ostream &os, set<T> v) { for (auto it = v.begin(); it != v.end(); it++) os << cond(it != v.begin()," ","") << *it; return os; }
template <typename F, typename S> ostream &operator<<(ostream &os, pair<F, S> v) { os << v.first << " " << v.second; return os; }

// General IO
template <typename... A> void in(A &...a) { ((cin >> a), ...); }
template <typename... A> void out(A... a) { ((cout << a << " "), ...); cout << endl; }
template <typename... A> void print(A... a) { ((cout << a), ...); }


int32_t main() {
    ios_base::sync_with_stdio(false); cin.tie(0); cout.tie(0);

    ll n, q, sum = 0; in(n, q);

    vector<ll> a(n+1, 0);
    vector<ll> d(n+2, 0);

    for (int i = 1; i <= n; i++) {
        in(a[i]);
    }

    while (q > 0) {
        int l, r; in(l, r);

        d[l] += 1;
        d[r+1] -= 1;

        --q;
    }

    // build prefix sum of deta vector
    ll acc = 0;
    for (int i = 1; i <= n; i++) {
        acc += d[i];
        d[i] = acc;
    }

    sort(a.begin()+1, a.end());
    sort(d.begin()+1, d.end()-1);

    for (int i = 1; i <= n; i++) {
        sum += a[i]*d[i];
    }

    out(sum);

    return 0;
}