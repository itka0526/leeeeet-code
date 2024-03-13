#include <bits/stdc++.h>

using namespace std;

int main()
{
    int k, m = 0;
    while (cin >> k && k != 0)
        if (k > m)
            m = k;
    cout << m << endl;
    return 0;
}