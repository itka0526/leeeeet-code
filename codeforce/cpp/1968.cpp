#include <bits/stdc++.h>

using namespace std;

void solveA() {
    int x;
    cin >> x;
    int val = 0, ans = 1;
    for (int y = x - 1; y > 0; --y) {
        int curr = gcd(x, y) + y;
        if (curr > val) {
            val = curr;
            ans = y;
        }
    }
    cout << ans << endl;
}

void solveB() {
    int n, m;
    cin >> n >> m;
    string a, b;
    cin >> a >> b;

    int k = 0;
    int j = 0;

    for (int i = 0; i < n; ++i) {
        while (j < m && b[j] != a[i]) {
            ++j;
        }
        if (j < m) {
            ++k;
            ++j;
        } else {
            break;
        }
    }
    std::cout << k << std::endl;
}

void solveC() {
    int n;
    cin >> n;
    vector<int> vec(n - 1);
    for (int i = 0; i < n - 1; i++) cin >> vec[i];
    vector<int> ans(n);
    ans[0] = 100000;
    for (int i = 1; i < n; i++) {
        ans[i] = ans[i - 1] + vec[i - 1];
    };
    for (auto i : ans) std::cout << i << ' ';
    cout << endl;
}

void solveD() {
    // TLE
    int64_t n, k, PB, PS;
    cin >> n >> k >> PB >> PS;
    vector<int64_t> p(n);
    vector<int64_t> a(n);
    for (int i = 0; i < n; i++) cin >> p[i];
    for (int i = 0; i < n; i++) cin >> a[i];
    int64_t s1, s2;
    s1 = s2 = 0;
    while (k > 0) {
        int64_t c1 = a[PB - 1], c2 = a[PS - 1];
        s1 += c1;
        s2 += c2;
        // Should the play~~
        cout << PB << " " << PS << endl;
        if (a[p[PB - 1] - 1] >= c1) {
            PS = p[PB - 1];
        }
        if (a[p[PS - 1] - 1] >= c2) {
            PS = p[PS - 1];
        }
        k--;
    }
    if (s1 > s2) {
        cout << "Bodya";
    } else if (s1 < s2) {
        cout << "Sasha";
    } else {
        cout << "Draw";
    }
    cout << endl;
}

int main() {
    int t;
    cin >> t;
    while (t-- > 0) {
        solveD();
    };
    return 0;
}