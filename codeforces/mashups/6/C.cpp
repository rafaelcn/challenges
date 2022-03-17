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
template <typename T>
istream &operator>>(istream &is, vector<T> &v)
{
    for (auto &a : v)
        is >> a;
    return is;
}
template <typename T>
ostream &operator<<(ostream &os, vector<T> v)
{
    loop(i, sz(v)) os << cond(i, " ", "") << v[i];
    return os;
}
template <typename T>
ostream &operator<<(ostream &os, set<T> v)
{
    for (auto it = v.begin(); it != v.end(); it++)
        os << cond(it != v.begin(), " ", "") << *it;
    return os;
}
template <typename F, typename S>
ostream &operator<<(ostream &os, pair<F, S> v)
{
    os << v.first << " " << v.second;
    return os;
}

// General IO
template <typename... A>
void in(A &...a) { ((cin >> a), ...); }
template <typename... A>
void out(A... a)
{
    ((cout << a << " "), ...);
    cout << endl;
}
template <typename... A>
void print(A... a) { ((cout << a), ...); }

int32_t main()
{
    ios_base::sync_with_stdio(false);
    cin.tie(0);
    cout.tie(0);

    ll n, m; in(n, m);

    ll moves = 0;

    queue<pair<int, int>> q;
    vector<bool> v(1e5 + 2, false);

    q.push({n, 0});

    while (!q.empty()) {
        tie(n, moves) = q.front();
        q.pop();

        //out(n, moves, m);

        if (n == m) {
            break;
        }

        if (n+1 <= 1e5 and !v[n+1]) {
            v[n+1] = true;
            q.push({n+1, moves+1});
        }

        if (n-1 > 0 and !v[n-1]) {
            v[n-1] = true;
            q.push({n-1, moves+1});
        }

        if (n*2 <= 1e5 and !v[n*2]) {
            v[n*2] = true;
            q.push({n*2, moves+1});
        }

        if (n%2 == 0 and n/2 > 0 and !v[n/2]) {
            v[n/2] = true;
            q.push({n/2, moves+1});
        }
    }

    out(moves);

    return 0;
}