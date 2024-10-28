#include <bits/stdc++.h>

using namespace std;

template<typename T>
class cpq : public priority_queue<T, vector<T>>
{
  public:
      bool remove(const T& value) {
        auto it = find(this->c.begin(), this->c.end(), value);
        if (it != this->c.end()) {
            this->c.erase(it);
            make_heap(this->c.begin(), this->c.end(), this->comp);
            return true;
       }
       else {
        return false;
       }
 }
};

int32_t main() {
    ios::sync_with_stdio(false);
    cin.tie();
    cout.tie();

    int t;
    stack<int> presents;
    cpq<int> weights;

    cin >> t;

    while (t > 0) {
        char c;
        int n;

        cin >> c >> n;

        if (c == 'A') {
            presents.push(n);
            weights.push(n);
        } else if (c == 'R') {
            int t = presents.top();
            presents.pop();

            weights.remove(t);
        } else if (c == 'V') {
            // show present
            if (weights.size() > 0) {
                cout << weights.top() << "\n";
            } else {
                cout << "0\n";
            }
        }

        t--;
    }

    return 0;
}