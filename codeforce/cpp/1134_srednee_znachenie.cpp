#include <bits/stdc++.h>

using namespace std;

int main()
{
    int n = 0, s = 0, k;
    while (cin >> k && k != 0)
    {
        n += 1;
        s += k;
    }
    cout << fixed << setprecision(6);
    cout << s * 1.0 / n << endl;
    return 0;
}