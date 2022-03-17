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

vector<string> split(string& s, string delimiter);
string join(vector<string>& s, string delimiter);

int main() {
    ll t; in(t);

    while (t > 0) {
        int n; in(n);

        if (n == 0) {
            t--;
            continue;
        }

        string students;
        string frequency;

        cin.ignore();

        getline(cin, students);
        getline(cin, frequency);

        auto sv = split(students, " ");
        auto fv = split(frequency, " ");

        vector<string> students_result;

        for (size_t i = 0; i < sv.size(); ++i) {
            float absent = 0.0;
            float size = fv[i].size();

            for (size_t j = 0; j < fv[i].size(); ++j) {
                if (fv[i][j] == 'A') {
                    absent++;
                } else if (fv[i][j] == 'M') {
                    size -= 1;
                }
            }

            float calc = absent/size;

            if (calc > 0.25) {
                students_result.push_back(sv[i]);
            }
        }

        cout << join(students_result, " ") << endl;
        t--;
    }


    return 0;
}

vector<string> split(string& s, string delim) {
    vector<string> tokens;

    size_t pos;
    string token;

    while ((pos = s.find(delim)) != string::npos) {
        token = s.substr(0, pos);
        s.erase(0, pos+1);

        tokens.push_back(token);
    }

    // if there's no ending delimiter we have to add the last word
    tokens.push_back(s);

    return tokens;
}

string join(vector<string>& s, string delim) {
    string r = "";

    if (s.size() == 0) {
        return "";
    } else if (s.size() == 1) {
        return s[0];
    }

    for (auto it = s.begin(); it < s.end(); ++it) {
        if (it == s.end()-1) {
            r += *it;
        } else {
            r += *it + delim;
        }
    }

    return r;
}