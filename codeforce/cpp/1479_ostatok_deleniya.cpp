#include <bits/stdc++.h>

using namespace std;

int main()
{
    long long a, b;
    cin >> a >> b;
    cout << (a % b + abs(b)) % abs(b);
    return 0;
}