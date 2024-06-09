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
    void dfs(int i, int j, vector<vector<char>> &grid, int n, int m, int *cnt)
    {
        if (grid[i][j] == '1')
        {
            grid[i][j] = '*';
            vector<pair<int, int>> dirs = {{1, 0}, {-1, 0}, {0, 1}, {0, -1}};
            for (pair<int, int> dir : dirs)
            {
                int ni = i + dir.first, nj = j + dir.second;
                if (ni >= 0 && ni < n && nj >= 0 && nj < m)
                    dfs(ni, nj, grid, n, m, cnt);
            }
        }
    };
    int numIslands(vector<vector<char>> &grid)
    {
        int n = grid.size();
        int m = grid[0].size();
        int cnt = 0;

        for (int i = 0; i < n; i++)
        {
            for (int j = 0; j < m; j++)
            {
                if (grid[i][j] == '1')
                {
                    dfs(i, j, grid, n, m, &cnt);
                    cnt++;
                }
            }
        }
        return cnt;
    }
};

int main()
{
    vector<vector<char>> g = {
        {'1', '1', '0', '0', '1'}, {'1', '1', '0', '0', '1'}, {'0', '0', '1', '0', '0'}, {'0', '0', '0', '1', '1'}};

    auto sol = new Solution();
    cout << sol->numIslands(g) << endl;
    return 0;
}