#include <bits/stdc++.h>

using namespace std;

void solveA() {
    int x, y;
    cin >> x >> y;
    if (x > y) {
        cout << y << " " << x << '\n';
    } else {
        cout << x << " " << y << '\n';
    }
}

void solveB() {
    string s;
    cin >> s;
    int i = 1, p = 0;
    if (s.size() <= 1) {
        cout << "NO"
             << "\n";
        return;
    }
    while (i < s.size()) {
        if (s[i] != s[p]) {
            swap(s[i], s[p]);
            cout << "YES"
                 << "\n"
                 << s << "\n";
            return;
        }
        p = i;
        i++;
    }
    cout << "NO"
         << "\n";
}

void solveC() {
    int a, b, c, d;
    cin >> a >> b >> c >> d;
    if (a > b) swap(a, b);
    if (c > d) swap(c, d);
    if ((a < c && c < b && b < d) || (c < a && a < d && d < b)) {
        cout << "YES\n";
    } else {
        cout << "NO\n";
    }
}

int main() {
    int q;
    cin >> q;
    while (q--) {
        solveC();
    }
    return 0;
}