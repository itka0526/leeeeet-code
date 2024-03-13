#include <bits/stdc++.h>

using namespace std;

int main()
{
    int a, b;
    cin >> a >> b;
    int u = a, v = b;
    while (v %= u)
        swap(u, v);
    cout << (a * b) / u << endl;
    return 0;
}