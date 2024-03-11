#include <bits/stdc++.h>

using namespace std;

int main()
{
    int64_t a, b, c, d;
    cin >> a >> b >> c >> d;
    int64_t x = -100;
    while (x <= 100)
    {
        if (a * (x * x * x) + b * (x * x) + (c * x) + d == 0)
            cout << x << " ";
        x++;
    }
    return 0;
}