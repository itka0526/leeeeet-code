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

class Solution
{
  public:
    int titleToNumber(string columnTitle)
    {
        int n = 0;
        for (char ch : columnTitle)
            n = n * 26 + (ch - 'A' + 1);
        return n;
    }
};

int main()
{
    auto s = new Solution();
    cout << s->titleToNumber("ZY") << nl;
    return 0;
}