#include <bits/stdc++.h>

using namespace std;

int main()
{
    int n, i = 1;
    cin >> n;
    while (n--)
    {
        int b;
        cin >> b;
        if (b <= 437)
        {
            cout << "Crash " << i << endl;
            return 0;
        }
        i++;
    }
    cout << "No crash";
    return 0;
}