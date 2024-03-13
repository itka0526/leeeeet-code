#include <bits/stdc++.h>

using namespace std;

int main()
{
    double x, y, k = 1;
    cin >> x >> y;
    while (x < y && fabs(x - y) > 1e-6)
    {
        x *= 1.15;
        k++;
    }
    cout << k << endl;
    return 0;
}