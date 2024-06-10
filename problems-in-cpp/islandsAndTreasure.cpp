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
    int W = -1;
    int TC = 0;
    int L = 2147483647;

    // Should return the closest distance
    void dfs(int i, int j, int n, int m, int dist, int &minDist, const vector<vector<int>> &grid,
             vector<vector<int>> &vis)
    {
        if (i < 0 || j < 0 || i >= n || j >= m || grid[i][j] == W || dist >= minDist || vis[i][j] == 1)
        {
            return;
        }
        if (grid[i][j] == TC)
        {
            minDist = min(minDist, dist);
        }
        vis[i][j] = 1;
        const vector<vector<int>> dirs = {{1, 0}, {-1, 0}, {0, 1}, {0, -1}};
        for (const vector<int> &dir : dirs)
        {
            dfs(i + dir[0], j + dir[1], n, m, dist + 1, minDist, grid, vis);
        }
        vis[i][j] = 0;
    }

    void islandsAndTreasure(vector<vector<int>> &grid)
    {
        int n = grid.size(), m = grid[0].size();

        for (int i = 0; i < n; i++)
        {
            for (int j = 0; j < m; j++)
            {
                if (grid[i][j] == L)
                {
                    vector<vector<int>> vis(n, vector<int>(m));
                    int minDist = L;
                    dfs(i, j, n, m, 0, minDist, grid, vis);
                    grid[i][j] = minDist;
                }
            }
        }
    };
};

int main()
{
    auto sol = new Solution();
    vector<vector<int>> g1 = {{2147483647, -1, 0, 2147483647},
                              {2147483647, 2147483647, 2147483647, -1},
                              {2147483647, -1, 2147483647, -1},
                              {0, -1, 2147483647, 2147483647}};
    vector<vector<int>> g2 = {{-1, 0, 2147483647}, {2147483647, 2147483647, -1}};
    vector<vector<int>> g3 = {
        {2147483647, 2147483647, 2147483647}, {2147483647, -1, 2147483647}, {0, 2147483647, 2147483647}};

    sol->islandsAndTreasure(g1);
    sol->islandsAndTreasure(g2);
    sol->islandsAndTreasure(g3);

    auto g = g3;
    for (int i = 0; i < g.size(); i++)
    {
        for (int j = 0; j < g[0].size(); j++)
        {
            cout << g[i][j] << ' ';
        }
        cout << '\n';
    }
    return 0;
}