#include <bits/stdc++.h>

using namespace std;

int main()
{
    int a, b, n;
    cin >> a >> b >> n;
    // Turn all money into the smallest possible currency
    b += a * 100;
    // Total
    b *= n;
    cout << b / 100 << " " << b % 100;
    return 0;
}