#include <bits/stdc++.h>

using namespace std;

int main()
{
    int k, e = 0;
    while (cin >> k && k != 0)
        if (k % 2 == 0)
            e++;
    cout << e << endl;
    return 0;
}