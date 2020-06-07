#include <bits/stdc++.h>
#include <cinttypes>

using namespace std;

bitset<32> to_bin(uint32_t n);
uint32_t to_dec(bitset<32> bin);

int main() {

    int s = 0;
    cin >> s;

    while (s) {
        uint32_t n = 0;
        cin >> n;

        bitset<32> binr = to_bin(n);

        cout << to_dec(binr.flip()) << "\n";
        s--;
    }


    return 0;
}

bitset<32> to_bin(uint32_t n) {
    size_t bit_s = n ? floor(log2(n))+1 : 0;

    bitset<32> bin(0);
    size_t i = 0;

    while (n > 0) {
        bin[i] = n%2;
        n = n/2;

        if (i > bit_s) {
            break;
        }
        else {
            i++;
        }
    }

    return bin;
}

uint32_t to_dec(bitset<32> bin) {
    uint32_t n = 0;
    int i = 31;

    while (i >= 0) {
        n += bin[i]*pow(2, i);
        i--;
    }

    return n;
}
