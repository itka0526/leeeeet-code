#include <bits/stdc++.h>

using namespace std;

void tsifra() {
    char n;
    cin >> n;
    cout << (n >= 48 && n <= 57 ? "Yes" : "No") << endl;
}

void verhnii_registr() {
    char ch;
    cin >> ch;
    if (ch >= 97 && ch <= 122)
        cout << (char)(ch + ('A' - 'a')) << endl;
    else
        cout << ch << endl;
}

void smena_registra() {
    char ch;
    cin >> ch;

    if (ch >= 97 && ch <= 122)
        cout << (char)(ch + ('A' - 'a')) << endl;
    else if (ch >= 65 && ch <= 90)
        cout << (char)(ch - ('A' - 'a')) << endl;
    else
        cout << ch << endl;
}

void kolichestvo_nulei() {
    string s;
    cin >> s;
    cout << count(s.begin(), s.end(), '0') << endl;
}

void udalenie_tsifr() {
    string s;
    cin >> s;
    string s1;
    for (auto &ch : s) {
        if (ch == '4' || ch == '7') {
            continue;
        }
        s1 += ch;
    }
    cout << s1 << endl;
}

void vstavka_symbolov() {
    string s;
    cin >> s;
    string s1;
    int idx = 0;
    for (auto &ch : s) {
        s1 += ch;
        if (idx + 1 >= s.length()) break;
        s1 += '#';
        idx += 1;
    }
    cout << s1 << endl;
}

void kruglyashi() {
    string s;
    cin >> s;
    int cnt = 0;
    for (auto &ch : s) {
        if (ch == '0' || ch == '6' || ch == '9') {
            cnt += 1;
        } else if (ch == '8') {
            cnt += 2;
        }
    }
    cout << cnt << endl;
}

bool isnumber(char ch) { return ch >= (char)'0' && ch <= (char)'9'; }

void paroli() {
    string s;
    cin >> s;
    int securityLevel = 0;
    bool containsUpper, containsLower, containsNumber;
    containsUpper = containsLower = containsNumber = false;
    if (s.length() >= 12) {
        securityLevel += 1;
        for (auto &ch : s) {
            if (isupper(ch) && !containsUpper) {
                securityLevel += 1;
                containsUpper = true;
            } else if (islower(ch) && !containsLower) {
                securityLevel += 1;
                containsLower = true;
            } else if (isnumber(ch) && !containsNumber) {
                securityLevel += 1;
                containsNumber = true;
            }
        }
    }
    cout << (securityLevel >= 4 ? "Yes" : "No") << endl;
}

int main() {
    // tsifra();
    // verhnii_registr();
    // smena_registra();
    // kolichestvo_nulei();
    // udalenie_tsifr();
    // vstavka_symbolov();
    // kruglyashi();`

    paroli();
    return 0;
}