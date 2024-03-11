#include <bits/stdc++.h>

using namespace std;

int main()
{
    ios::sync_with_stdio(false);
    cin.tie(NULL);
    cout.tie(NULL);

    int s1 = 0, s2 = 0, n = 4;
    while (n > 0)
    {
        int a, b;
        cin >> a >> b;
        s1 += a;
        s2 += b;
        n--;
    }
    if (s1 == s2)
        cout << "DRAW";
    else if (s1 < s2)
        cout << 2;
    else
        cout << 1;
    return 0;
}