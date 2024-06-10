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
    int E = 0, F = 1, R = 2;
    int orangesRotting(vector<vector<int>> &grid)
    {
        int minElapsed = 0, remaining = 0;
        int n = grid.size(), m = grid[0].size();
        queue<pair<int, int>> q;

        for (int i = 0; i < n; i++)
        {
            for (int j = 0; j < m; j++)
            {
                if (grid[i][j] == F)
                {
                    remaining++;
                }
                else if (grid[i][j] == R)
                {
                    q.push({i, j});
                }
            }
        }

        const vector<vector<int>> &dirs = {{1, 0}, {-1, 0}, {0, 1}, {0, -1}};
        // Each tick?
        while (remaining > 0 && !q.empty())
        {
            int l = q.size();
            // Queue = [[R, R, R], [R, R, R, R, R, R]] and so on!
            // Clever trick!
            for (int k = 0; k < l; k++)
            // END OF TRICK!
            {
                pair<int, int> curr = q.front();
                int i = curr.first, j = curr.second;
                q.pop();
                for (const vector<int> &dir : dirs)
                {
                    if (i + dir[0] >= 0 && j + dir[1] >= 0 && i + dir[0] < n && j + dir[1] < m)
                    {
                        if (grid[i + dir[0]][j + dir[1]] == F)
                        {
                            grid[i + dir[0]][j + dir[1]] = R;
                            q.push({i + dir[0], j + dir[1]});
                            remaining--;
                        }
                    }
                }
            }
            minElapsed++;
        }
        return remaining == 0 ? minElapsed : -1;
    }
};

int main()
{
    vector<vector<int>> g1 = {{1, 1, 0},
                              {
                                  0,
                                  1,
                                  1,
                              },
                              {
                                  0,
                                  1,
                                  2,
                              }};
    vector<vector<int>> g2 = {{2, 1, 1}, {1, 1, 0}, {0, 1, 1}};
    vector<vector<int>> g = g2;
    for (int i = 0; i < g.size(); i++)
    {
        for (int j = 0; j < g[0].size(); j++)
        {
            cout << g[i][j] << ' ';
        };
        cout << '\n';
    }
    auto sol = new Solution();
    cout << sol->orangesRotting(g) << '\n';
    return 0;
}