#include <bits/stdc++.h>

using namespace std;

int main()
{
    int n;
    cin >> n;
    int a[n], b[n];

    for (int i = 0; i < n; i++)
        cin >> a[i];

    for (int i = 0; i < n; i++)
        cin >> b[i];

    int h = 0;
    for (int i = 1; i < n; i++)
    {
        if (a[h] * b[h] < a[i] * b[i])
        {
            h = i;
        }
    }
    cout << h + 1;
    return 0;
}