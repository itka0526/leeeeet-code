// C.Left and Right Houses

#include <bits/stdc++.h>

using namespace std;

int main()
{
    int q;
    cin >> q;
    while (q--)
    {
        int n;
        cin >> n;
        vector<int> a(n);
        vector<int> suf(n + 1, 0);
        for (auto &x : a)
            cin >> x;
        for (int i = n - 1; i >= 0; i--)
        {
            suf[i] = suf[i + 1] + (a[i] ? 1 : 0);
        }
        int cnt = 0;
    }
    return 0;
}