#include <iostream>
#include <string>
#include <stack>
#include <map>

inline bool check_size(const std::stack<char>& s) {
    if (s.size() == 0) {
        return false;
    }

    return true;
}

bool match(const std::string& s) {
    bool r = false;

    std::stack<char> matcher;

    std::map<char, char> pairs = {
        {'}', '{'},
        {']', '['},
        {')', '('},
    };

    for (const auto& e: s) {
        switch (e) {
        case '{':
        case '[':
        case '(':
            matcher.push(e);
            break;
        case '}':
        case ']':
        case ')':
            bool r = check_size(matcher);

            if (r && matcher.top() == pairs.at(e)) {
                matcher.pop();
            } else {
                return false;
            }
            break;
        }
    }


    if (matcher.size() == 0) {
        r = true;
    } else {
        r = false;
    }

    return r;
}

int main() {
    using namespace std;

    int t;
    cin >> t;

    while (t > 0) {
        string s;
        cin >> s;

        bool b = match(s);

        if (b) {
            cout << "YES\n";
        } else {
            cout << "NO\n";
        }

        --t;
    }

    return 0;
}