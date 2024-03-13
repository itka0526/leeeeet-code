#include <bits/stdc++.h>

using namespace std;

int main()
{
    int k, mx, mn;
    int i = 1;
    cin >> mn >> mx;
    while (cin >> k)
    {
        if (i % 2 == 0 && mx < k)
            mx = k;
        else if (i % 2 == 1 && mn > k)
            mn = k;
        i += 1;
    }
    cout << mn + mx << endl;
    return 0;
}