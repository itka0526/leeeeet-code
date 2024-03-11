#include <bits/stdc++.h>

using namespace std;

int main()
{
    int n;
    cin >> n;
    int ma = 0, mi = 0;
    bool didRun = false;
    for (int i = 0; i < n; i++)
    {
        int a, g;
        cin >> a >> g;
        if (g == 1 && a > ma)
        {
            ma = a;
            mi = i;
            didRun = true;
        }
    };
    cout << (didRun ? mi + 1 : -1) << endl;
    return 0;
}