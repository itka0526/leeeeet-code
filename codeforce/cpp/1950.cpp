#include <bits/stdc++.h>

using namespace std;

void solveA() {
    int a, b, c;
    cin >> a >> b >> c;
    if (a < b && b < c)
        cout << "STAIR";
    else if (b > a && b > c)
        cout << "PEAK";
    else
        cout << "NONE";
    cout << endl;
}

void solveB() {
    int n;
    cin >> n;
    bool change = true;
    for (int i = 0; i < n * 2; i++) {
        for (int j = 0; j < n; j++) {
            if (j % 2 == 0 && change) {
                cout << "##";
            } else if (j % 2 == 1 && !change) {
                cout << "##";
            } else {
                cout << "..";
            }
        }
        cout << endl;
        if ((i + 1) % 2 == 0) change = !change;
    }
}

void solveC() {
    int h, m;
    scanf("%d:%d", &h, &m);
    printf("%02d:%02d %s\n", h % 12 == 0 ? 12 : h % 12, m, h >= 12 && h < 24 ? "PM" : "AM");
}

bool checkOnly01(const string &s) {
    return s.size() == count_if(s.begin(), s.end(), [&](const char &c) { return c == '0' || c == '1'; });
}

void solveD() {
    string n;
    cin >> n;
    // First check if the n contains only 0s and 1s
    cout << (checkOnly01(n) ? "YES" : "NO") << endl;
}

void solveE() {
    int n;
    string s;
    cin >> n >> s;
    string s1;
    for (int i = 0; i < s.size(); i++) {
        char c = s[i];
        int duplIdx = s.find(s1, i);
        if (!s1.empty() && duplIdx == string::npos) {
            break;
        };
        s1 += c;
    }
    cout << s1 << " " << s << endl;
}

int main() {
    int q;
    cin >> q;
    while (q--) solveE();
    return 0;
}