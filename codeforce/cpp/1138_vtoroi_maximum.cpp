#include <bits/stdc++.h>

using namespace std;

int main()
{
    int k;
    vector<int> a;
    while (cin >> k && k != 0)
        a.push_back(k);
    sort(a.begin(), a.end(), greater<int>());
    cout << a[1] << endl;
    return 0;
}