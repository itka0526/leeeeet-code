#include <bits/stdc++.h>

using namespace std;

int main()
{
    double a, b, c;
    cin >> a >> b >> c;

    if (a == 0 && b == 0 && c == 0)
    {
        cout << -1;
        return 0;
    }
    if (a == 0 && b == 0 && c != 0)
    {
        cout << 0;
        return 0;
    }

    double d = b * b - 4 * a * c;
    cout << fixed << setprecision(6);
    if (a == 0 && b != 0)
    {
        cout << 1 << endl << (-c / b);
        return 0;
    }
    if (d < 0)
    {
        cout << 0;
    }
    else if (d == 0)
    {
        cout << 1 << endl << (-b + sqrt(d)) / (2 * a) << endl;
    }
    else
    {
        cout << 2 << endl << (-b - sqrt(d)) / (2 * a) << endl << (-b + sqrt(d)) / (2 * a) << endl;
    }

    return 0;
}