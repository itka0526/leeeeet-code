#include <bits/stdc++.h>

using namespace std;

int main()
{
    int x, ans;
    cin >> x;
    ans = to_string(x - 1).size() > to_string(x + 1).size() ? x + 1 : x - 1;
    cout << ans << endl;
    return 0;
}