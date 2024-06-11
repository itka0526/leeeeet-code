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
    void dfs(int i, int j, int n, int m, const vector<vector<int>> &heights, int prev, set<pair<int, int>> &vis)
    {
        if (i < 0 || j < 0 || i >= n || j >= m || vis.count({i, j}))
            return;
        int curr = heights[i][j];
        if (prev > curr)
            return;
        vis.insert({i, j});
        vector<vector<int>> dirs = {{1, 0}, {-1, 0}, {0, 1}, {0, -1}};
        for (auto dir : dirs)
            dfs(i + dir[0], j + dir[1], n, m, heights, curr, vis);
    }
    vector<vector<int>> pacificAtlantic(vector<vector<int>> &heights)
    {
        vector<vector<int>> ans;
        int n = heights.size(), m = heights[0].size();
        set<pair<int, int>> atl, pac;

        // Determine if the cell is pacific
        for (int i = 0; i < n; i++)
            dfs(i, 0, n, m, heights, heights[i][0], pac);
        for (int j = 0; j < m; j++)
            dfs(0, j, n, m, heights, heights[0][j], pac);
        // Determine if the cell is atlantic
        for (int i = 0; i < n; i++)
            dfs(i, m - 1, n, m, heights, heights[i][m - 1], atl);
        for (int j = 0; j < m; j++)
            dfs(n - 1, j, n, m, heights, heights[n - 1][j], atl);

        for (int i = 0; i < n; i++)
        {
            for (int j = 0; j < m; j++)
            {
                if (pac.count({i, j}) && atl.count({i, j}))
                {
                    ans.push_back({i, j});
                }
            }
        }
        return ans;
    }
};

int main()
{
    vector<vector<int>> h1 = {{4, 2, 7, 3, 4}, {7, 4, 6, 4, 7}, {6, 3, 5, 3, 6}};
    auto sol = new Solution();

    auto h = h1;
    auto ans = sol->pacificAtlantic(h);
    for (int i = 0; i < ans.size(); i++)
    {
        for (int num : ans[i])
        {
            cout << num << ' ';
        }
        cout << '\n';
    }
    return 0;
}