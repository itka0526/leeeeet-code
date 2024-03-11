#include <bits/stdc++.h>

using namespace std;

int main()
{
    int n, m;
    cin >> n >> m;
    // total_number_of_quadrants = n * m
    // crossovers = (n + 1) * m + (m + 1) * n
    if (n < m)
    {
        int shelf = n;
        n = m;
        m = shelf;
    }
    if (n == 1 || m == 1)
        cout << n * m * 4;
    else
    {
        if (n % 2 == 1 && m % 2 == 1)
            cout << (m + 2) * n + m * (n + 2) - 2;
        else
            cout << (m + 2) * n + m * (n + 2);
    }
    return 0;
}