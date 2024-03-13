#include <bits/stdc++.h>

using namespace std;

int main()
{
    int c, p, s;
    cin >> p;
    s = p;
    while (cin >> c && !(c == 0 && p == 0))
    {
        s += c;
        p = c;
    }
    cout << s << endl;
    return 0;
}