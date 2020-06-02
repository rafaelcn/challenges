#include <iostream>
#include <string>
#include <vector>
#include <cassert>

using namespace std;

vector<string> split(string& s, string delimiter);
string join(vector<string>& s, string delimiter);

int main() {
    int n = 0;

    // number of entries
    cin >> n;

    while (n > 0) {
        // I don't care about this
        int n_students = 0;
        cin >> n_students;

        string students;
        string check;

        getline(cin, students);
        getline(cin, check);

        auto students_v = split(students, " ");
        auto check_v = split(check, " ");

        vector<string> students_result;

        for (int i = 0; i < students_v.size(); ++i) {
            int absent = 0;

            for (int j = 0; j < check_v[i].size(); ++j) {
                if (check_v[i][j] == 'A') {
                    absent++;
                }
            }

            assert(check_v[i].size() != 0);
            float calc = absent/check_v[i].size();

            if (calc < 0.75) {
                students_result.push_back(students_v[i]);
            }
        }

        cout << join(students_result, " ");
        cout << "\n";
        n--;
    }


    return 0;
}

vector<string> split(string& s, string delimiter) {
    vector<string> split;

    size_t pos;
    string token;

    while ((pos = s.find(delimiter)) != string::npos) {
        token = s.substr(0, pos);
        s.erase(0, pos);

        split.push_back(token);
    }

    // if there's no ending delimiter we have to add the last word
    split.push_back(s);

    return split;
}

string join(vector<string>& s, string delimiter) {
    string r = "";

    for(auto it = s.begin(); it <= s.end(); ++it) {
        if (it == s.end()) {
            r += *it;
        } else {
            r += *it + delimiter;
        }
    }

    return r;
}