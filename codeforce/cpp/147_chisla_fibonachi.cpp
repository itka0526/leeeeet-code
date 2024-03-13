#include <bits/stdc++.h>

using namespace std;

int main()
{
    int n;
    cin >> n;
    int t1 = 0, t2 = 1;
    int t3 = t1 + t2;

    while (n--)
    {
        t1 = t2;
        t2 = t3;
        t3 = t1 + t2;
    }
    cout << t1 << endl;
    return 0;
}