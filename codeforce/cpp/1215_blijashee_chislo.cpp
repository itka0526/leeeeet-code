#include <bits/stdc++.h>

using namespace std;

int main()
{
    int n;
    cin >> n;
    int a[n];

    for (int i = 0; i < n; i++)
    {
        cin >> a[i];
    }

    int x;
    cin >> x;

    // Keep track of the closest
    int d = a[0];
    for (int i = 0; i < n; i++)
    {
        if (abs(a[i] - x) < abs(d - x) || abs(a[i] - x) == abs(d - x) && a[i] < d)
        {
            d = a[i];
        }
    }
    cout << d;
    return 0;
}
