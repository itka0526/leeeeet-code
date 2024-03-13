#include <bits/stdc++.h>

using namespace std;

int main()
{
    int k, m, i = 1;
    cin >> m;
    while (cin >> k && k != 0)
    {
        if (k > m)
        {
            m = k;
            i = 1;
        }
        else if (m == k)
        {
            i += 1;
        }
    };
    cout << i << endl;
    return 0;
}