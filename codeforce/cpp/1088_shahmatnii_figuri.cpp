#include <bits/stdc++.h>

using namespace std;

int main()
{
    ios::sync_with_stdio(false);
    cin.tie(NULL);
    cout.tie(NULL);

    char l1, cn1, l2, cn2;
    cin.get(l1);
    cin.get(cn1);
    cin.get();
    cin.get(l2);
    cin.get(cn2);

    string a = "ABCDEFGH";
    int x1 = a.find(l1) + 1;
    int y1 = cn1 - '0';

    int x2 = a.find(l2) + 1;
    int y2 = cn2 - '0';

    bool s = true;

    if ((x1 != x2 && y1 == y2) || (x1 == x2 && y1 != y2))
    {
        cout << "Rook" << endl;
        s = false;
    }

    if (abs(x2 - x1) == abs(y2 - y1))
    {
        cout << "Bishop" << endl;
        s = false;
    }

    int d1 = abs(x2 - x1);
    int d2 = abs(y2 - y1);

    bool d3 = (d1 == 1 && y2 == y1) || (d2 == 1 && x2 == x1) || (d1 == 1 && d2 == 1);
    if (!d3 && ((abs(x2 - x1) == 2 && abs(y2 - y1) == 1) || (abs(x2 - x1) == 1 && abs(y2 - y1) == 2)))
    {
        cout << "Knight" << endl;
        s = false;
    }
    if ((x1 != x2 && y1 == y2) || (x1 == x2 && y1 != y2) || abs(x2 - x1) == abs(y2 - y1))
    {
        cout << "Queen" << endl;
        s = false;
    }
    if ((d1 == 1 && y2 == y1) || (d2 == 1 && x2 == x1) || (d1 == 1 && d2 == 1))
    {
        cout << "King" << endl;
        s = false;
    }
    if ((x1 == x2 && y1 != 1) && (y1 == 2 && y1 + 2 == y2 || y1 + 1 == y2))
    {
        cout << "Pawn" << endl;
        s = false;
    }
    if (s)
        cout << "Nobody";
    return 0;
}