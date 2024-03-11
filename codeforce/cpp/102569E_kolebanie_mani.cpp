#include <bits/stdc++.h>

using namespace std;

int main()
{
    int n;
    cin >> n;
    int64_t b = 0, m = 0;
    for (int i = 0, a; i < n; i++)
    {
        cin >> a;
        b += a;
        if (b < m)
            m = b;
    }
    cout << abs(m) << endl;
    return 0;
}