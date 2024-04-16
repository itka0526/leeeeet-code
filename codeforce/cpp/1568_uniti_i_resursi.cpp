#include <bits/stdc++.h>

using namespace std;

int main() {
    int n, m;
    cin >> n >> m;
    vector<int> resources(n);
    for (int i = 0; i < n; i++) cin >> resources[i];
    while (m--) {
        bool canAfford = 1;
        for (int i = 0; i < n; i++) {
            int reqResource;
            cin >> reqResource;
            if (reqResource > resources[i]) canAfford = 0;
        }
        cout << canAfford << " ";
    }
    return 0;
}