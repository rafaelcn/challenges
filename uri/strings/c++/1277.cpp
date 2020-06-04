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
        // I don't care about the number of students given to the input as I
        // read the entire line.
        int n_students = 0;
        cin >> n_students;

        string students;
        string check;

        // ignore new line from the last read buffer.
        cin.ignore();

        getline(cin, students);
        getline(cin, check);

        auto students_v = split(students, " ");
        auto check_v = split(check, " ");

        vector<string> students_result;

        for (size_t i = 0; i < students_v.size(); ++i) {
            float absent = 0.0;
            float size = check_v[i].size();

            for (size_t j = 0; j < check_v[i].size(); ++j) {
                if (check_v[i][j] == 'A') {
                    absent++;
                } else if (check_v[i][j] == 'M') {
                    size -= 1;
                }
            }

            assert(check_v[i].size() != 0);
            float calc = absent/size;

            if (calc >= 0.25) {
                students_result.push_back(students_v[i]);
            }
        }

        cout << join(students_result, " ");
        cout << endl;
        --n;
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