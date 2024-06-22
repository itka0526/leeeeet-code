#include <bits/stdc++.h>
#define nl '\n'
#define ll long long
#define YES cout << "YES" << nl
#define NO cout << "NO" << nl
#define vi vector<int>
#define vll vector<long long>

using namespace std;

ostream &operator<<(ostream &os, const vector<pair<int, int>> &v)
{
    os << "[ ";
    for (const pair<int, int> &p : v)
        os << "(" << p.first << ", " << p.second << ")" << " ";
    os << "]\n";
    return os;
}

template <typename T> ostream &operator<<(ostream &os, const vector<T> &v)
{
    os << "[ ";
    for (const auto &item : v)
        os << item << " ";
    os << "]\n";
    return os;
}

void notQuiteLatinSquare()
{
    vector<string> m(3);
    for (int i = 0; i < 3; i++)
    {
        cin >> m[i];
    }

    int d = 'C' + 'A' + 'B';

    for (int i = 0; i < 3; i++)
    {
        for (int j = 0; j < 3; j++)
        {
            if (m[i][j] == '?')
            {
                for (int k = 0; k < 3; k++)
                {
                    if (k == j)
                        continue;
                    d -= m[i][k];
                }
                cout << char(d) << nl;
                return;
            }
        }
    }
}

void maximumMultipleSum()
{
    int n;
    cin >> n;
    int x = 2;
    pair<int, int> ans = {0, 0};
    for (; x <= n; x++)
    {
        int localMax = 0;
        for (int k = 1; k * x <= n; k++)
        {
            localMax += k * x;
        }
        if (localMax > ans.second)
        {
            ans = {x, localMax};
        }
    }
    cout << ans.first << nl;
    // int n; cin >> n;
    // cout << (n == 3 ? 3 : 2) << nl;
}

void goodKid()
{
    int n;
    cin >> n;
    vector<int> a(n);
    for (int &num : a)
        cin >> num;
    vector<int>::iterator x = min_element(a.begin(), a.end());
    (*x)++;
    int ans = accumulate(a.begin(), a.end(), 1, [](int acc, int x) { return acc * x; });
    cout << ans << nl;
}

void tenWordsOfWisdom()
{
    int n;
    cin >> n;
    pair<int, int> winner = {-1, 0};
    for (int i = 0; i < n; i++)
    {
        int a, b;
        cin >> a >> b;
        if (a <= 10 && b > winner.second)
            winner = {i + 1, b};
    }
    cout << winner.first << nl;
}

void blankSpace()
{
    int n;
    cin >> n;
    vector<int> a(n);
    for (int &num : a)
        cin >> num;
    int ls = 0, lls = 0;
    for (int i = 0; i < n; i++)
    {
        if (a[i] == 0)
            lls++;
        else
            lls = 0;
        ls = max(lls, ls);
    }
    cout << ls << nl;
}

void grabTheCandies()
{
    int n;
    cin >> n;
    auto cmp = [](const int &c1, const int &c2) {
        if (c1 % 2 == 0)
        {
            if (c2 % 2)
            {
                return false;
            }
            else
            {
                return c1 < c2;
            }
        }
        else if (c2 % 2 == 0)
        {
            return true;
        }
        else
        {
            return c1 < c2;
        }
    };
    priority_queue<int, vector<int>, decltype(cmp)> pq(cmp);
    for (int i = 0, t; i < n; i++)
    {
        cin >> t;
        pq.push(t);
    }
    int a = 0, b = 0;
    while (!pq.empty())
    {
        int curr = pq.top();
        pq.pop();
        if (curr % 2 == 0)
            a += curr;
        else
            b += curr;
        if (a <= b)
        {
            NO;
            return;
        }
    }
    YES;
    // int n; cin >> n;
    // int a = 0, b = 0;
    // for (int i = 0, t; i < n; i++)
    //   cin >> t; if (t % 2 == 0) a += t; else b += t;
    // a > b ? YES : NO;
}

int main()
{
    int t;
    cin >> t;
    while (t--)
    {
        // notQuiteLatinSquare();
        // maximumMultipleSum();
        // goodKid();
        // tenWordsOfWisdom();
        // blankSpace();
        grabTheCandies();
    }
    return 0;
}