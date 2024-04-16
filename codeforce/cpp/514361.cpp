#include <bits/stdc++.h>

using namespace std;

void solveA() {
    int n;
    cin >> n;
    vector<short> vec;
    for (int i = 0; i < n; i++) {
        char ch;
        cin >> ch;
        vec.push_back(ch);
    };
    sort(vec.begin(), vec.end(), greater<int>());
    cout << vec[0] - 'a' + 1 << endl;
}

template <int id>
void kek(int n, auto& cnt) {
    for (int i = 0; i < n; i++) {
        std::string str;
        cin >> str;
        std::get<id>(cnt[str]) = true;
    }
}

int sum(const int items[3]) { return items[0] + items[1] + items[2]; }

void solveB() {
    int numberOfWords;
    cin >> numberOfWords;

    map<string, int[3]> record;
    int numberOfPlayers = 3;

    for (int i = 0; i < numberOfPlayers; i++) {
        for (int j = 0; j < numberOfWords; j++) {
            string currentWord;
            cin >> currentWord;
            record[currentWord][i]++;
        }
    }

    int a, b, c;
    a = b = c = 0;

    for (auto& [key, value] : record) {
        int totalScores = sum(value);
        for (int playerId = 0; playerId < 3; playerId++) {
            int score = value[playerId];
            if (score && totalScores != 3) {
                int incrementBy;
                if (totalScores == 2) {
                    incrementBy = 1;
                } else {
                    incrementBy = 3;
                }

                if (playerId == 0)
                    a += incrementBy;
                else if (playerId == 1)
                    b += incrementBy;
                else
                    c += incrementBy;
            }
        }
    }

    cout << a << " " << b << " " << c << endl;
}

void solveC() {
    int n;
    cin >> n;
}

void solveD() {
    int n;
    cin >> n;
    map<string, int> cnt;
    vector<string> v(n);
    for (int i = 0; i < n; i++) {
        string s;
        cin >> s;
        v[i] = s;
        ++cnt[s];
    }
    string ans(n, '0');
    for (int j = 0; j < n; j++) {
        string str = v[j];
        for (int i = 0; i + 1 < str.size(); i++) {
            if (cnt[str.substr(0, i + 1)] && cnt[str.substr(i + 1)]) {
                ans[j] = '1';
                break;
            }
        }
    }
    cout << endl;
}

int main() {
    int t;
    cin >> t;
    while (t--) solveB();
    return 0;
}