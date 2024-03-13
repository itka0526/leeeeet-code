#include <bits/stdc++.h>

using namespace std;

int main()
{
    int k, p, s = 0;
    cin >> p;
    while (cin >> k && k != 0)
    {
        if (k > p)
            s++;
        p = k;
    }
    cout << s << endl;
    return 0;
}