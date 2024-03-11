#include <bits/stdc++.h>

using namespace std;

int main()
{
    int x1, y1, x2, y2;
    cin >> x1 >> y1 >> x2 >> y2;

    int d1 = abs(x2 - x1);
    int d2 = abs(y2 - y1);

    if ((d1 == 1 && y2 == y1) || (d2 == 1 && x2 == x1) || (d1 == 1 && d2 == 1))
        cout << "YES";
    else
        cout << "NO";
    return 0;
}