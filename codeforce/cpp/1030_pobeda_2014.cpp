
#include <bits/stdc++.h>

using namespace std;

int main()
{
    long double a1, a2, a3, a4;
    cin >> a1 >> a2 >> a3 >> a4;
    cout << (int64_t)floor(sqrt(min(a1, a2) + min(a3, a4))) << endl;
    return 0;
}