#include <bits/stdc++.h>

using namespace std;

int main()
{
    int n, s1;
    int total = 0;
    cin >> n;

    int a[n];
    for (int i = 0; i < n; i++)
        cin >> a[i];

    cin >> s1;
    for (int s2 : a)
        if (s1 == s2)
            total++;

    cout << total;
    return 0;
}