#include <bits/stdc++.h>

using namespace std;

int main()
{
    int n;
    cin >> n;
    int i = 0;
    while ((1 << i) < n)
        i++;
    cout << ((1 << i) == n ? "YES" : "NO");
    return 0;
}