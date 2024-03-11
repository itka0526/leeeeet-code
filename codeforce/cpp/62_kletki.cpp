
#include <bits/stdc++.h>

using namespace std;

int main()
{
    ios::sync_with_stdio(false);
    cin.tie(NULL);
    cout.tie(NULL);

    char l, cn;
    cin.get(l);
    cin.get(cn);

    string a = "ABCDEFGH";
    int x = a.find(l) + 1;
    int y = cn - '0';
    if ((x % 2 == 1 && y % 2 == 1) || (x % 2 == 0 && y % 2 == 0))
        cout << "BLACK";
    else
        cout << "WHITE";
    return 0;
}