#include <bits/stdc++.h>

using ll = long long;

int solve(std::stack<int>& x, std::stack<int> y, ll sum) {
    int moves = 0;
    int result = 0;

    do {
        int tx = 0;
        int ty = 0;

        if (x.size() > 0) {
            tx = x.top();
        }

        if (y.size() > 0) {
            ty = y.top();
        }

        if (tx < ty) {
            if (tx + result >= sum) {
                break;
            }

            moves += 1;
            result += tx;
            x.pop();
        } else {
            if (ty + result >= sum) {
                break;
            }

            moves += 1;
            result += ty;
            y.pop();
        }
    } while ((x.size() > 0 || y.size() > 0));

    return moves;
}

int main() {
    using namespace std;

    int n;
    cin >> n;

    int result = 0;

    while (n > 0) {
        int nx, ny;
        ll sum;

        cin >> nx >> ny >> sum;

        stack<int> x;
        stack<int> y;
        // I have to create this as the input given by hackerrank is totally
        // messed up as they expect the last inserted number to be at the
        // beginning of the stack.
        vector<int> xv(nx);
        vector<int> yv(ny);

        // filling xv
        for (size_t i = 0; i < nx; i++) {
            cin >> xv[i];
        }
        // filling yv
        for (size_t i = 0; i < ny; i++) {
            cin >> yv[i];
        }

        // inverting array and inserting as hackerrank demands.
        for (size_t i = xv.size()-1; i >= 0; i--) {
            x.push(xv[i]);
        }
        for (size_t i = yv.size()-1; i >= 0; i--) {
            y.push(yv[i]);
        }

        cout << solve(x, y, sum) << "\n";

        --n;
    }

    return 0;
}