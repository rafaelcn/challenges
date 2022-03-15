#include <iostream>
#include <vector>

using namespace std;

string lcs(string word1, string word2);

int main() {
    /*while (true) {
        int n;
        // number of words
        cin >> n;

        if (n == 0) {
            break;
        }

        vector<string> words;
        while (n > 0) {
            string word;
            cin >> word;

            words.push_back(word);

            --n;
        }

        if (words.size() == 1) {
            cout << words[0].size() << "\n";
        } else {
            string s = "";
            string m = "";

            // O(m*n)
            for (size_t i = 1; i <= words.size()-1; i++) {
                for (size_t j = 0; j < i; j++) {
                    // O(m*n)
                    m = lcs(words[i], words[j]);

                    if (m.size() > s.size()) {
                        s = m;
                    }
                }
            }

            cout << s << " " << s.size() << "\n";
        }
    }*/

    cout << lcs("rafael", "afa");
    cout << lcs("afa", "rafael");
    cout << lcs("rafael", "el");
    cout << lcs("afa", "afa");
    cout << lcs("afaafaafaafaafaafaafaafaafaafa", "afaaf");
}

string lcs(string word1, string word2) {
    //
    int max_len = 0;
    //
    int last = word1.size();
    // lookup table of the two words
    vector<vector<int>> lookup;

    // initializing container
    lookup.resize(word1.size());
    for (auto &e : lookup) {
        e.resize(word2.size());
    }

    for (size_t w1 = 1; w1 <= word1.size(); w1++) {
        for (size_t w2 = 1; w2 <= word2.size(); w2++) {
            if (word1[w1-1] == word2[w2-1]) {
                lookup[w1][w2] = lookup[w1-1][w2-1] + 1;

                if (lookup[w1][w2] > max_len) {
                    max_len = lookup[w1][w2];
                    last = w1;
                }
            }
        }
    }

    auto common = word1.substr(last - max_len, last);

    return common;
}