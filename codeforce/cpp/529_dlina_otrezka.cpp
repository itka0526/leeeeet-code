#include <bits/stdc++.h>

using namespace std;

int main()
{
    ios::sync_with_stdio(false);
    cin.tie(NULL);
    cout.tie(NULL);
    int x1, y1, x2, y2;
    cin >> x1 >> y1 >> x2 >> y2;
    int a = abs(x2 - x1);
    int b = abs(y2 - y1);
    double c = sqrt(a * a + b * b);
    cout << fixed << c << setprecision(5);
    return 0;
}