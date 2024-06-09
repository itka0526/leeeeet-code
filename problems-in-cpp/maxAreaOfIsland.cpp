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
    int dfs(int i, int j, int n, int m, vector<vector<int>> &g)
    {
        if (i < 0 || j < 0 || i >= n || j >= m || g[i][j] != 1)
            return 0;
        g[i][j] = 5;
        return 1 + dfs(i + 1, j, n, m, g) + dfs(i - 1, j, n, m, g) + dfs(i, j + 1, n, m, g) + dfs(i, j - 1, n, m, g);
    }

    int maxAreaOfIsland(vector<vector<int>> &grid)
    {
        int n = grid.size(), m = grid[0].size();
        int maxArea = 0;
        for (int i = 0; i < n; i++)
        {
            for (int j = 0; j < m; j++)
            {
                if (grid[i][j] == 1)
                {
                    maxArea = max(maxArea, dfs(i, j, n, m, grid));
                }
            }
        }
        return maxArea;
    }
};

int main()
{
    vector<vector<int>> g = {{0, 1, 1, 0, 1},
                             {
                                 1,
                                 0,
                                 1,
                                 0,
                             },
                             {
                                 0,
                                 1,
                                 1,
                                 0,
                             },
                             {
                                 0,
                                 1,
                                 0,
                                 0,
                             }};
    auto sol = new Solution();
    cout << "Solution: " << sol->maxAreaOfIsland(g) << endl;
    return 0;
}