#include <bits/stdc++.h>

using namespace std;

int main()
{
    int n, c3 = 0;
    cin >> n;
    while (n % 5 != 0)
    {
        n -= 3;
        c3++;
    }
    cout << n / 5 << " " << c3;
    return 0;
}