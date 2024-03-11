#include <bits/stdc++.h>

using namespace std;

int main()
{
    int64_t n;
    cin >> n;
    int64_t i = 0;
    while ((1 << i) <= n)
    {
        cout << (1 << i) << " ";
        i++;
    }
    return 0;
}