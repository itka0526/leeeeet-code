#include <bits/stdc++.h>

using namespace std;

int main()
{
    int x1, y1, x2, y2;
    int ax, ay, bx, by;
    cin >> x1 >> y1 >> x2 >> y2 >> ax >> ay;
    if (x1 == x2)
    {
        bx = 2 * x1 - ax;
        by = ay;
    }
    if (y1 == y2)
    {
        bx = ax;
        by = 2 * y1 - ay;
    }
    cout << bx << " " << by;
    return 0;
}