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
    void dfs(int i, int j, vector<vector<char>> &b)
    {
        vector<vector<int>> dirs = {{1, 0}, {-1, 0}, {0, 1}, {0, -1}};
        if (b[i][j] == 'O')
        {
            b[i][j] = '*';
            for (auto dir : dirs)
            {
                int ni = i + dir[0], nj = j + dir[1];
                if (ni >= 0 && nj >= 0 && ni < b.size() && nj < b[0].size())
                {
                    dfs(ni, nj, b);
                }
            }
        }
    }
    void solve(vector<vector<char>> &board)
    {
        int n = board.size(), m = board[0].size();
        // Left, right borders
        for (int i = 0; i < n; i++)
        {
            dfs(i, 0, board);
            dfs(i, m - 1, board);
        }
        // Top, bottom borders
        for (int j = 0; j < m; j++)
        {
            dfs(0, j, board);
            dfs(n - 1, j, board);
        }
        for (int i = 0; i < n; i++)
        {
            for (int j = 0; j < m; j++)
            {
                if (board[i][j] == '*')
                {
                    board[i][j] = 'O';
                }
                else
                {
                    board[i][j] = 'X';
                }
            }
        }
    }
};

int main()
{
    vector<vector<char>> g = {{'X', 'X', 'X', 'X'}, {'X', 'O', 'O', 'X'}, {'X', 'O', 'O', 'X'}, {'X', 'X', 'X', 'O'}};
    auto s = new Solution();
    s->solve(g);

    cout << '\n';
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