#include <bits/stdc++.h>

using namespace std;

int main()
{
    int n, k;
    cin >> n >> k;
    cout << k / n << " " << k % n << " ";
    if (k % n != 0)
        cout << n - k % n;
    else
        cout << 0;
    return 0;
}