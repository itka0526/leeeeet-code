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
    bool hasCycle(int child, int parent, map<int, set<int>> &m, set<int> &vis)
    {
        // Has cycle
        if (vis.count(child))
        {
            return true;
        }
        vis.insert(child);
        for (int curr : m[child])
        {
            // Must not be the previous node because our graph is bi-directional
            if (child != parent && hasCycle(curr, child, m, vis))
            {
                return true;
            }
        }
        return false;
    }
    bool validTree(int n, vector<vector<int>> &edges)
    {
        map<int, set<int>> m;
        for (vector<int> e : edges)
        {
            m[e[0]].insert(e[1]);
            m[e[1]].insert(e[0]);
        }
        for (pair<int, set<int>> p : m)
        {
            cout << p.first << ": ";
            for (int e : m[p.first])
            {
                cout << e << " ";
            }
            cout << "\n";
        }
        // Connected and has no cycles.
        set<int> vis;
        if (hasCycle(0, -1, m, vis))
        {
            return false;
        };
        cout << "FUCK: ";
        for (int i : vis)
            cout << i << " ";
        cout << '\n';
        return vis.size() == m.size();
    }
};

int main()
{
    auto s = new Solution();
    vector<vector<int>> e1 = {{0, 1}, {0, 2}, {0, 3}, {1, 4}};
    vector<vector<int>> e2 = {{0, 1}, {1, 2}, {2, 3}, {1, 3}, {1, 4}};
    cout << (s->validTree(5, e1) ? "YES" : "NO") << '\n';
    cout << (s->validTree(5, e2) ? "YES" : "NO") << '\n';
    return 0;
}