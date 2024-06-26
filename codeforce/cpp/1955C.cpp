#include <bits/stdc++.h>

using namespace std;

int main() {
    int q;
    cin >> q;
    while (q--) {
        int64_t n, k;
        cin >> n >> k;
        deque<int> a(n);
        for (auto& x : a) cin >> x;
        int cnt = 0;
        while (a.size() > 1 && k) {
            int64_t mn = min(a.front(), a.back());
            if (k < 2 * mn) {
                a.front() -= k / 2 + k % 2;
                a.back() -= k / 2;
                k = 0;
            } else {
                a.front() -= mn;
                a.back() -= mn;
                k -= 2 * mn;
            }
            if (a.front() == 0) {
                a.pop_front();
                cnt++;
            }
            if (a.back() == 0) {
                a.pop_back();
                cnt++;
            }
        }
        if (a.size() && a.front() <= k) cnt++;
        cout << cnt << '\n';
    }
}

int main0() {
    // C.Inhabitant of the Deep Sea
    // Two pointer problem
    int q;
    cin >> q;
    while (q > 0) {
        int n, k;
        cin >> n >> k;
        vector<int> a(n);
        for (auto& x : a) cin >> x;
        int left = 0;
        int right = n - 1;
        int cnt = 0;
        bool lr = true;
        while (k > 0) {
            if (left == right) {
                // Last ship
                if (k >= a[left]) {
                    cnt++;
                }
                break;
            } else if (lr) {
                a[left]--;
                if (a[left] == 0) {
                    left++;
                    cnt++;
                }
            } else {
                a[right]--;
                if (a[right] == 0) {
                    right--;
                    cnt++;
                }
            }
            lr = !lr;
            k--;
        }
        cout << cnt << '\n';
        q--;
    }
    return 0;
}