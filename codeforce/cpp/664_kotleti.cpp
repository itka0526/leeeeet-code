#include <bits/stdc++.h>

using namespace std;

int main()
{
    int k, m, n;
    cin >> k >> m >> n;

    // Naive approach
    // kotleti nado s obeih storon 2 * m
    // n kotlet
    // ceil(n / k)
    // int t = (n / k) + ((n % k) != 0);
    // int s = t * 2 * m;
    // cout << s;

    if (n < k)
        cout << 2 * m;
    else if (n % k == 0)
        cout << n / k * 2 * m;
    else
    {
        // Floor it down
        int ans = (n / k - 1) * 2 * m;
        // Leftovers
        int l = n % k;
        if (l <= k / 2)
            ans += 3 * m;
        else
            ans += 4 * m;
        cout << ans;
    }
    return 0;
}