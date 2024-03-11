#include <bits/stdc++.h>

using namespace std;

int main()
{
    ios::sync_with_stdio(false);
    cin.tie(NULL);
    cout.tie(NULL);

    int n;
    cin >> n;
    n / 1000 == n % 10 && n % 1000 / 100 == n % 100 / 10 ? cout << "YES" : cout << "NO";
    return 0;
}