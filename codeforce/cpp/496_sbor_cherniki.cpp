#include <bits/stdc++.h>

using namespace std;

int main()
{
    int n;
    cin >> n;
    int a[n];
    int m = 0;

    for (int i = 0; i < n; i++)
    {
        cin >> a[i];
        m = a[0] + a[n - 1] + ((a[1] > a[n - 2]) ? a[1] : a[n - 2]);
    }

    int i = 2;
    while (i < n)
    {
        int s = a[i] + a[i - 1] + a[i - 2];
        if (s > m)
        {
            m = s;
        }
        i++;
    }

    cout << m << "\n";
    return 0;
}