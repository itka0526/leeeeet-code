#include <iostream>

using namespace std;

int main() {
    ios::sync_with_stdio(false);
    cin.tie(NULL);
    cout.tie(NULL);

    int n, q;
    cin >> n;
    long long b[n + 1];
    b[0] = 0;
    for (int i = 0; i < n; i++) {
        long long t;
        cin >> t;
        b[i + 1] = t + b[i];
    };
    cin >> q;
    for (int i = 0; i < q; i++) {
        int n1, n2;
        cin >> n1 >> n2;
        if (n1 - 1 < 0) {
            cout << b[n2] - b[n1 - 1] << "\n";
        } else {
            cout << b[n2] - b[n1 - 1] << "\n";
        }
    }
}