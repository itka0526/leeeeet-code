#include <bits/stdc++.h>

using namespace std;

int main()
{
    int n;
    cin >> n;
    int i = 2;
    while (i < n && n % i != 0)
        i++;
    cout << i << endl;
    return 0;
}