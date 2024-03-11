#include <bits/stdc++.h>

using namespace std;

int main()
{
    ios::sync_with_stdio(false);
    cin.tie(NULL);
    cout.tie(NULL);
    int n;
    cin >> n;
    if (n > 1)
        if (n % 2 == 0)
            cout << n / 2;
        else
            cout << n;
    else
        cout << 0;
    return 0;
}