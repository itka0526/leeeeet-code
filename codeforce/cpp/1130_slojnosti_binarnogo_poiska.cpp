#include <bits/stdc++.h>

using namespace std;

int main()
{
    int64_t n;
    cin >> n;
    int k = 0;
    while ((n + 2) / 2 != 1)
    {
        n = n / 2 + (n % 2 != 0 ? 1 : 0);
        k += 1;
    };
    cout << k << endl;
}