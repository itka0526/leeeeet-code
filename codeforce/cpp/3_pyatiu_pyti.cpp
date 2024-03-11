#include <bits/stdc++.h>

using namespace std;

int main()
{
    int n;
    cin >> n;
    int a = n / 10;
    if (a * (a + 1) == 0)
        cout << 25;
    else
        cout << a * (a + 1) << 25;
    return 0;
}