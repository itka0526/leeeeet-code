#include <bits/stdc++.h>

using namespace std;

int main()
{
    int n, x, y;
    int g, l;
    cin >> n >> x >> y;
    // Making the x as the fastest printer
    if (x > y)
        swap(x, y);

    g = gcd(x, y); // Greatest common divisible number
    l = x * y / g; // Least common multiple

    int nl = l / x + l / y; // TODO: Mystery...

    int t;
    t = x; // First page
    n--;

    t += n / nl * l; // Adding the time when both printers are working
    n %= nl;         // Subtracting the times when both printers were working, so we don't (TLE)

    int t1 = 0, t2 = 0;
    while (n > 0)
    {
        t++;
        t1++;
        t2++;
        if (t1 == x)
        {
            n--;
            t1 = 0;
        }
        if (t2 == y)
        {
            n--;
            t2 = 0;
        }
    }
    cout << t << endl;
    return 0;
}