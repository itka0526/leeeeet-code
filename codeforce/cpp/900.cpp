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

void manhattanCircle()
{
    int n, m;
    cin >> n >> m;
    vector<string> g(n);
    // i-th, count of #
    pair<int, int> ic = {0, 0};
    for (int i = 0; i < n; i++)
    {
        cin >> g[i];
        int curr = count(g[i].begin(), g[i].end(), '#');
        if (ic.second < curr)
            ic = {i, curr};
    }
    for (int j = 0; j < m; j++)
    {
        if (g[ic.first][j] == '#')
        {
            cout << ic.first + 1 << " " << j + ic.second / 2 + 1 << nl;
            return;
        }
    }
}

int cantor(int a, int b)
{
    return (a + b + 1) * (a + b) / 2 + b;
}

int myHash(int a, int b, int c)
{
    return cantor(a, cantor(b, c));
}

// DP with memoization
void ties(int p1, int p2, int p3, int curr, int &maxTies, unordered_map<int, int> &memo)
{
    int k = myHash(p1 + 1, p2 + 1, p3 + 1);
    if (memo.count(k))
    {
        maxTies = max(maxTies, memo[k]);
        return;
    }
    memo[k] = curr;
    maxTies = max(maxTies, curr);
    // 3! = 6 possible outputs
    // p1 and p2, p2 and p1
    if (p1 > 0 && p2 > 0)
        ties(p1 - 1, p2 - 1, p3, curr + 1, maxTies, memo);
    // p1 and p3, p3 and p1
    if (p1 > 0 && p3 > 0)
        ties(p1 - 1, p2, p3 - 1, curr + 1, maxTies, memo);
    // p2 and p3, p3 and p2
    if (p2 > 0 && p3 > 0)
        ties(p1, p2 - 1, p3 - 1, curr + 1, maxTies, memo);
}

void chessForThree()
{
    // Winner - 2, loser - 0, tie - 1
    // p1 <= p2 <= p3
    int p1, p2, p3;
    cin >> p1 >> p2 >> p3;
    // Output the maximum amount of games that ended in tie. If not possible return -1.
    int ans = -1;
    if ((p1 + p2 + p3) % 2 == 1)
    {
        cout << ans << nl;
    }
    else
    {
        unordered_map<int, int> memo;
        ties(p1, p2, p3, 0, ans, memo);
        cout << ans << nl;
    }
}

void coinGames()
{
    int n;
    cin >> n;
    string s;
    cin >> s;
    int cntU = count(s.begin(), s.end(), 'U');
    cntU % 2 ? YES : NO;
}

int main()
{
    int t;
    cin >> t;
    while (t--)
    {
        // manhattanCircle();
        // chessForThree();
        coinGames();
    }
    return 0;
}