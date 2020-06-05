#include <iostream>
#include <vector>
#include <cmath>

#define printd(x) printf("%d\n", (x))
#define print(x) printf("%.15f\n", (x))

int main() {
    float n;

    std::vector<float> currency = {100.00, 50.00, 20.00, 10.00, 5.00, 2.00};
    std::vector<float> coins = {1.00, 0.50, 0.25, 0.10, 0.05, 0.01};

    std::vector<int> currency_qtd;
    std::vector<int> coins_qtd;

    std::cin >> n;

    for (auto& c: currency) {
        float qtd = std::floor(n/c);

        n -= c*qtd;

        currency_qtd.push_back(qtd);
    }

    for (auto& c: coins) {
        float qtd = std::floor(n/c);

        n -= c*qtd;

        coins_qtd.push_back(qtd);
    }

    std::cout << "NOTAS:\n";
    for (size_t i = 0; i < currency.size(); i++) {
        std::printf("%d nota(s) de R$ %.2f\n", currency_qtd[i], currency[i]);
    }

    std::cout << "MOEDAS:\n";
    for (size_t i = 0; i < coins.size(); i++) {
        std::printf("%d moeda(s) de R$ %.2f\n", coins_qtd[i], coins[i]);
    }

    return 0;
}