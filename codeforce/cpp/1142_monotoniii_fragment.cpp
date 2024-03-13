#include <bits/stdc++.h>

using namespace std;

int main()
{
    int c, p, a = 1, b = 1, sa = 1, sb = 1;
    cin >> p;

    while (p)
    {
        cin >> c;
        if (c > p)
            a++;
        else
        {
            if (a > sa)
                sa = a;
            a = 1;
        }

        if (c < p)
            b++;
        else
        {
            if (b > sb)
                sb = b;
            b = 1;
        }
        p = c;
    }
    cout << max(sa, sb);
    return 0;
}