#include <bits/stdc++.h>

using namespace std;

int main()
{
    int h1, m1, s1;
    int h2, m2, s2;
    cin >> h1 >> m1 >> s1;
    cin >> h2 >> m2 >> s2;
    cout << s2 - s1 + (m2 - m1) * 60 + (h2 - h1) * 60 * 60;
    return 0;
}