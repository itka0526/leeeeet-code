#include <bits/stdc++.h>

using namespace std;

int main()
{
    int x, p, y;
    cin >> x >> p >> y;
    int years = 0;
    x = x * 100;
    y = y * 100;
    p = p + 100;
    while (x < y)
    {
        x = (x * p) / 100;
        years += 1;
    }
    cout << years << endl;
    return 0;
}