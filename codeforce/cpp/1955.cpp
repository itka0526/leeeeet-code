#include <bits/stdc++.h>

using namespace std;

void solveA() {
    int a, b, n;
    cin >> n >> a >> b;
    if (a * 2 < b) {
        cout << n * a << endl;
    } else {
        cout << n / 2 * b + n % 2 * a << endl;
    }
}

void solveB() {
    int64_t n, c, d;
    cin >> n >> c >> d;

    vector<vector<int64_t>> a(n, vector<int64_t>(n));
    vector<vector<int64_t>> prog(n, vector<int64_t>(n));
    for (int i = 0; i < n; i++) {
        for (int j = 0; j < n; j++) {
            cin >> a[i][j];
        }
    }

    // a(bottom, curr) = a(curr, curr) + c;
    // a(curr, right) = a(curr, curr) + d;

    prog[0][0] = a[0][0];

    for (int i = 0; i < n; i++) {
        for (int j = 0; j < n; j++) {
            // a(bottom, curr) = a(curr, curr) + c;
            // a(curr, right) = a(curr, curr) + d;
            if (i + 1 < n) prog[i + 1][j] = prog[i][j] + c;
            if (j + 1 < n) prog[i][j + 1] = prog[i][j] + d;
        }
    }
    cout << "Original: " << endl;
    for (int i = 0; i < n; i++) {
        for (int j = 0; j < n; j++) {
            cout << a[i][j] << ' ';
        }
        cout << endl;
    }
    cout << "Progressive: " << endl;
    for (int i = 0; i < n; i++) {
        for (int j = 0; j < n; j++) {
            cout << prog[i][j] << ' ';
        }
        cout << endl;
    }
    cout << endl;
}

int main() {
    int t;
    cin >> t;
    while (t--) solveB();
    return 0;
}