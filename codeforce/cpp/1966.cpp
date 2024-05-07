#include <bits/stdc++.h>

using namespace std;

// A. Обмен картами
void solveA() {
    int n, k;
    cin >> n >> k;
    vector<int> c(n);
    for (int i = 0; i < n; i++) cin >> c[i];
}

// B. Покраска прямоугольника
void solveB() {
    int n, m;
    cin >> n >> m;
    vector<string> grid(n);

    for (int i = 0; i < n; i++) {
        cin >> grid[i];
    }
}

// C. Всё о Ниме
void solveC() {
    int n;
    std::cin >> n;
    std::vector<int> v(n);
    for (int i = 0; i < n; i++) cin >> v[i];

    int min_pile = *std::min_element(v.begin(), v.end());
    int count_min_pile = std::count(v.begin(), v.end(), min_pile);

    cout << min_pile << " " << count_min_pile << endl;
}

// D. Отсутствующая сумма подпоследовательности
void solveD() {
    int n, k;
    cin >> n >> k;
}

// E. Складывание полоски
void solveE() {
    int n;
    string s;
    cin >> n >> s;
}

// F. Отсутствующая сумма подмассива
void solveF() {
    int n;
    cin >> n;
    vector<int> sums((n * (n + 1)) / 2 - 1);
    for (int i = 0; i < (n * (n + 1)) / 2 - 1; i++) {
        cin >> sums[i];
    }
}

int main() {
    int t;
    cin >> t;
    while (t--) solveF();
    return 0;
}