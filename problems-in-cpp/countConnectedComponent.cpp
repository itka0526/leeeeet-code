#include <bits/stdc++.h>
#define nl '\n'
#define ll long long
#define YES cout << "YES" << nl
#define NO cout << "NO" << nl
#define vi vector<int>
#define vll vector<long long>
using namespace std;

class Solution
{
  public:
    void dfs(int curr, map<int, set<int>> &m, set<int> &v)
    {
        if (v.count(curr))
            return;
        v.insert(curr);
        set<int> neighbors = m[curr];
        for (int next : neighbors)
            dfs(next, m, v);
    }
    int countComponents(int n, const vector<vector<int>> &edges)
    {
        if (edges.size() == 0)
        {
            return n;
        }
        map<int, set<int>> m;
        for (vector<int> e : edges)
        {
            m[e[0]].insert(e[1]);
            m[e[1]].insert(e[0]);
        }
        set<int> vis;
        int cnt = 0;
        for (pair<int, set<int>> p : m)
        {
            int curr = p.first;
            if (vis.count(curr))
            {
                continue;
            }
            dfs(curr, m, vis);
            cnt++;
        }
        return cnt;
    }
};

int main()
{

    return 0;
}