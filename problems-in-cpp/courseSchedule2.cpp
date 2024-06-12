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
    bool dfs(int i, map<int, set<int>> &m, set<int> &vis, set<int> &cl, vector<int> &ans)
    {
        if (cl.count(i))
            return false;
        if (vis.count(i))
            return true;
        // Detect cycle
        cl.insert(i);
        for (int j : m[i])
            if (!dfs(j, m, vis, cl, ans))
                return false;
        cl.erase(i);
        vis.insert(i);
        ans.push_back(i);
        return true;
    }
    vector<int> findOrder(int numCourses, vector<vector<int>> &prerequisites)
    {
        vector<int> ans;
        map<int, set<int>> m;
        for (const vector<int> &p : prerequisites)
        {
            int i = p[0], j = p[1];
            if (!m.count(j))
                m[j] = {};
            m[i].insert(j);
        }
        set<int> vis, cl;
        for (int i = 0; i < numCourses; i++)
            if (!dfs(i, m, vis, cl, ans))
                return {};
        return ans;
    }
};

int main()
{
    return 0;
}