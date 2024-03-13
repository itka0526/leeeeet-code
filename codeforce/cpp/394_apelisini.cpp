#include <bits/stdc++.h>

using namespace std;

int main()
{
    int64_t a, b;
    cin >> a >> b;
    int64_t u = a, v = b;
    while (v %= u)
        swap(u, v);
    cout << ((a * b) / u) / b << endl;
    return 0;
}