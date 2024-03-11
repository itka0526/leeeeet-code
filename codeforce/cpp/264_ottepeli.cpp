#include <bits/stdc++.h>

using namespace std;

int main()
{
    int n;
    int m = 0;
    int ma = 0;
    cin >> n;
    while (n--)
    {
        int c;
        cin >> c;
        if (c > 0)
        {
            m += 1;
            if (m > ma)
                ma = m;
        }
        else
        {
            m = 0;
        }
    }
    cout << ma << endl;
    return 0;
}