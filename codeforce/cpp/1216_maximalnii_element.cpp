#include <bits/stdc++.h>

using namespace std;

int main()
{
    int n;
    cin >> n;
    int a[n + 1];
    for (int i = 1; i <= n; i++)
    {
        cin >> a[i];
    }
    int l, r, mi;
    cin >> l >> r;
    mi = l;
    for (int j = l + 1; j <= r; j++)
    {
        if (a[j] > a[mi])
        {
            mi = j;
        }
    }
    cout << a[mi] << " " << mi;
    return 0;
}