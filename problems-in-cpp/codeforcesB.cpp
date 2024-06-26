#include <bits/stdc++.h>
#define nl '\n'
#define ll long long
#define YES cout << "YES" << nl
#define NO cout << "NO" << nl
#define vi vector<int>
#define vll vector<long long>

using namespace std;

template <typename T> ostream &operator<<(ostream &os, pair<T, T> p)
{
    os << "(" << p.first << ", " << p.second << ")";
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

template <typename T, typename Sequence> ostream &operator<<(ostream &os, queue<T, Sequence> &q)
{
    vector<T> v(q.size());
    int i = 0;
    while (!q.empty())
    {
        v[i++] = q.front();
        q.pop();
    }
    os << "Q: " << v;
    for (auto e : v)
        q.push(e);
    return os;
}

template <typename T, typename Container, typename Compare>
ostream &operator<<(ostream &os, priority_queue<T, Container, Compare> &pq)
{
    vector<T> v(pq.size());
    int i = 0;
    while (!pq.empty())
    {
        v[i++] = pq.top();
        pq.pop();
    }
    os << "PQ: " << v;
    for (auto e : v)
        pq.push(e);
    return os;
}

template <typename T, typename Compare, typename Allocator>
ostream &operator<<(ostream &os, set<T, Compare, Allocator> &s)
{
    os << "S: { ";
    for (const auto &item : s)
        os << item << " ";
    os << "}\n";
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

void increasing()
{
    int n;
    cin >> n;
    priority_queue<int, vector<int>, greater<int>> pq;
    for (int i = 0, tmp; i < n; i++)
    {
        cin >> tmp;
        pq.push(tmp);
    }
    int prev = pq.top();
    pq.pop();
    while (!pq.empty())
    {
        int cur = pq.top();
        pq.pop();
        if (cur <= prev)
        {
            NO;
            return;
        }
        prev = cur;
    }
    YES;
}

void colorBlindness()
{
    int n;
    cin >> n;
    string r1, r2;
    cin >> r1 >> r2;
    for (int i = 0; i < n; i++)
    {
        // R from B, R from G
        if ((r1[i] == 'R' && (r2[i] == 'B' || r2[i] == 'G')) || (r2[i] == 'R' && (r1[i] == 'B' || r1[i] == 'G')))
        {
            NO;
            return;
        }
    }
    YES;
}

void equalCandies()
{
    int n;
    cin >> n;
    int ans = 0;
    vector<int> candies(n);
    for (int &num : candies)
        cin >> num;
    int minCandy = *min_element(candies.begin(), candies.end());
    for (int i = 0; i < n; i++)
        ans += candies[i] - minCandy;
    cout << ans << nl;
}

void triple()
{
    int n, x;
    cin >> n;
    map<int, int> m;
    while (n--)
    {
        cin >> x;
        m[x]++;
        if (m[x] >= 3)
        {
            cout << x << nl;
            cin.ignore(numeric_limits<streamsize>::max(), '\n');
            return;
        }
    }
    cout << -1 << nl;
}

void samePartySummands()
{
    int n, k;
    cin >> n >> k;
    int n1 = n - (k - 1);
    if (n1 > 0 && n1 % 2 == 1)
    {
        YES;
        for (int i = 0; i < k - 1; i++)
            cout << "1 ";
        cout << n1 << nl;
        return;
    }
    int n2 = n - 2 * (k - 1);
    if (n2 > 0 && n2 % 2 == 0)
    {
        YES;
        for (int i = 0; i < k - 1; i++)
            cout << "2 ";
        cout << n2 << nl;
        return;
    }
    NO;
}

void vlandAndShapes()
{
    int n;
    cin >> n;
    assert(n >= 2);
    vector<string> g(n);
    for (string &s : g)
        cin >> s;
    set<int> cntOnes;
    for (int i = 0; i < n; i++)
        cntOnes.insert(count(g[i].begin(), g[i].end(), '1'));
    cntOnes.erase(0);
    cout << (cntOnes.size() == 1 ? "SQUARE" : "TRIANGLE") << nl;
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
        // grabTheCandies();
        // increasing();
        // colorBlindness();
        // equalCandies();
        // triple();
        // samePartySummands();
        vlandAndShapes();
    }
    return 0;
}