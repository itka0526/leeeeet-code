#include <bits/stdc++.h>

using namespace std;

int main()
{
    int n;
    cin >> n;
    for (int i = 0; i < n; i++)
    {
        int t;
        cin >> t;
        string ans = "";
        int k = 0;
        if (t / 10000 % 10 > 0)
        {
            ans.append(to_string(t / 10000 % 10 * 10000) + " ");
            k += 1;
        }
        if (t / 1000 % 10 > 0)
        {
            ans.append(to_string(t / 1000 % 10 * 1000) + " ");
            k += 1;
        }
        if (t / 100 % 10 > 0)
        {
            ans.append(to_string(t / 100 % 10 * 100) + " ");
            k += 1;
        }
        if (t / 10 % 10 > 0)
        {
            ans.append(to_string(t / 10 % 10 * 10) + " ");
            k += 1;
        }
        if (t / 1 % 10 > 0)
        {
            ans.append(to_string(t / 1 % 10 * 1) + " ");
            k += 1;
        }
        cout << k << "\n" << ans << endl;
    }
    return 0;
}