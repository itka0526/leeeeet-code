#include <bits/stdc++.h>

using namespace std;

int main()
{
    ios::sync_with_stdio(false);
    cin.tie(NULL);
    cout.tie(NULL);

    int a, b, c;
    cin >> a >> b >> c;
    if (a < 94 || a > 727 || b < 94 || b > 727 || c < 94 || c > 727)
        cout << "Error";
    else
        cout << max(a, max(b, c));
    return 0;
}