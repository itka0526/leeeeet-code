#include <bits/stdc++.h>
#define nl '\n'
#define ll long long
#define YES cout << "YES" << nl
#define NO cout << "NO" << nl
#define vi vector<int>
#define vll vector<long long>
using namespace std;

int main()
{
    int n, q;
    cin >> n >> q;
    vector<pair<int, int>> c(n);
    for (int i = 0; i < n; i++)
    {
        int a, b;
        cin >> a >> b;
        c[i] = {a, b};
    }

    sort(c.begin(), c.end(), [](const pair<int, int> &x, const pair<int, int> &y) { return x.second < y.second; });

    int time = 0;
    int tip = 0;

    for (int i = 0; i < n; i++)
    {
        time += c[i].second;
        tip += c[i].first - time;
    }
    cout << tip << nl;

    while (q--)
    {
        int i, A, B;
        cin >> i >> A >> B;
        --i;
        c[i] = {A, B};
        sort(c.begin(), c.end(), [](const pair<int, int> &x, const pair<int, int> &y) { return x.second < y.second; });
        tip = 0;
        time = 0;
        for (int j = 0; j < n; j++)
        {
            time += c[j].second;
            tip += c[j].first - time;
        }
        cout << tip << nl;
    }

    return 0;
}