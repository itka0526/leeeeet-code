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
    bool dfs1(int i, map<int, set<int>> &m, set<int> &vis)
    {
        if (vis.count(i))
        {
            return false;
        }
        // If the set is empty then its a truthy
        if (m[i].empty())
        {
            return true;
        }
        // If not for each set item recurse till it finds them and return true
        vis.insert(i);
        for (int j : m[i])
        {
            if (!dfs1(j, m, vis))
            {
                return false;
            }
        }
        vis.erase(i);
        m[i].clear();
        return true;
    }
    bool canFinishDFS(int numCourses, vector<vector<int>> &prerequisites)
    {
        map<int, set<int>> m;
        for (vector<int> p : prerequisites)
        {
            if (!m.count(p[0]))
                m[p[0]] = set<int>();
            m[p[0]].insert(p[1]);
        }

        set<int> vis;
        for (pair<int, set<int>> p : m)
        {
            if (!dfs1(p.first, m, vis))
            {
                return false;
            };
        }
        return true;
    }

    bool dfs2(int i, map<int, set<int>> &m, set<int> &vis)
    {
        if (vis.count(i))
        {
            return false;
        }
        if (m[i].empty())
        {
            return true;
        }
        vis.insert(i);
        for (int j : m[i])
        {
            if (!dfs2(j, m, vis))
            {
                return false;
            }
        }
        vis.erase(i);
        m[i].clear();
        return true;
    }

    bool canFinish2(int numCourses, vector<vector<int>> &prerequisites)
    {
        map<int, set<int>> m;
        set<int> vis;

        for (vector<int> p : prerequisites)
        {
            int i = p[0], j = p[1];
            if (!m.count(i))
                m[i] = {};
            m[i].insert(j);
        }

        for (pair<int, set<int>> p : m)
        {
            if (!dfs2(p.first, m, vis))
            {
                return false;
            };
        }

        for (auto p : m)
        {
            cout << p.first << ": ";
            for (auto g : p.second)
            {
                cout << g << " ";
            }
            cout << "\n";
        }

        return true;
    }

    bool dfs(int i, map<int, set<int>> &m, set<int> &v)
    {
        if (v.count(i))
            return false;
        if (m[i].empty())
            return true;
        v.insert(i);
        for (int j : m[i])
        {
            if (!dfs(j, m, v))
            {
                return false;
            }
        }
        v.erase(i);
        m[i].clear();
        return true;
    }

    bool canFinish(int numCourses, vector<vector<int>> &prerequisites)
    {
        map<int, set<int>> m;
        for (const vector<int> &p : prerequisites)
        {
            int i = p[0], j = p[1];
            if (!m.count(j))
                m[j] = {};
            m[i].insert(j);
        }
        set<int> vis;
        for (pair<int, set<int>> p : m)
        {
            if (!dfs(p.first, m, vis))
            {
                return false;
            }
        }
        for (auto p : m)
        {
            cout << p.first << ": ";
            for (auto g : p.second)
            {
                cout << g << " ";
            }
            cout << "\n";
        }
        return true;
    }

    bool dfs4(int i, map<int, set<int>> &m, set<int> &v)
    {
        if (v.count(i))
            return false;
        if (m[i].empty())
            return true;
        v.insert(i);
        for (int j : m[i])
        {
            if (!dfs(j, m, v))
            {
                return false;
            }
        }
        v.erase(i);
        m[i].clear();
        return true;
    }

    bool canFinish4(int numCourses, vector<vector<int>> &prerequisites)
    {
        map<int, set<int>> m;
        for (vector<int> p : prerequisites)
        {
            int i = p[0], j = p[1];
            if (!m.count(j))
                m[j] = {};
            m[i].insert(j);
        }
        set<int> v;
        for (pair<int, set<int>> p : m)
        {
            if (!dfs4(p.first, m, v))
            {
                return false;
            }
        }
        return true;
    }
};

int main()
{
    int n1 = 2;
    vector<vector<int>> p1 = {{0, 1}, {1, 0}};

    int n2 = 5;
    vector<vector<int>> p2 = {{1, 4}, {2, 4}, {3, 1}, {3, 2}};

    auto n = n1;
    auto p = p1;
    auto s = new Solution();
    cout << "Can finish: \n" << s->canFinish(n, p) << endl;

    n = n2;
    p = p2;
    cout << "Can finish: \n" << s->canFinish(n, p) << endl;
    return 0;
}