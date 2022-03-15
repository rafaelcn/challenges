#include <iostream>
#include <vector>
#include <cmath>

#define printd(x) printf("%d\n", (x))
#define print(x) printf("%.15f\n", (x))

void count(int& n, const std::vector<float>& v, std::vector<int>& r) {
    for (auto& c: v) {
        float qtd = std::floor(n/c);

        n -= c*qtd;

        r.push_back(qtd);
    }
}

int main() {
    std::string n;

    std::vector<float> currency = {100.00, 50.00, 20.00, 10.00, 5.00, 2.00};
    // there's a bit of a hacking to display this correctly
    std::vector<float> coins = {100, 50, 25, 10, 5, 1};

    std::vector<int> currency_qtd;
    std::vector<int> coins_qtd;

    std::cin >> n;

    int curre = std::atoi(n.substr(0, n.find(".")).c_str());
    int cents = 0;

    if (n.find(".") != std::string::npos) {
         cents = std::atoi(n.substr(n.find(".")+1, n.size()-1).c_str());
    }

    count(curre, currency, currency_qtd);
    count(cents, coins, coins_qtd);

    if (curre < 10) {
        coins_qtd[0] = curre;
    }

    std::cout << "NOTAS:\n";
    for (size_t i = 0; i < currency.size(); i++) {
        std::printf("%d nota(s) de R$ %.2f\n", currency_qtd[i], currency[i]);
    }

    std::cout << "MOEDAS:\n";
    for (size_t i = 0; i < coins.size(); i++) {
       if (coins[i] == 100) {
            std::printf("%d moeda(s) de R$ 1.00\n", coins_qtd[i]);
       } else if (coins[i] <= 50 && coins[i] > 5) {
            std::printf("%d moeda(s) de R$ 0.%.0f\n", coins_qtd[i], coins[i]);
       } else {
           std::printf("%d moeda(s) de R$ 0.0%.0f\n", coins_qtd[i], coins[i]);
       }
    }

    return 0;
}