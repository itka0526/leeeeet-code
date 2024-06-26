#include <bits/stdc++.h>
#define nl '\n'
#define ll long long
#define YES cout << "YES" << nl
#define NO cout << "NO" << nl
#define vi vector<int>
#define vll vector<long long>

using namespace std;

ostream &operator<<(ostream &os, const vector<pair<int, int>> &v)
{
    os << "[ ";
    for (const pair<int, int> &p : v)
        os << "(" << p.first << ", " << p.second << ")" << " ";
    os << "]\n";
    return os;
}

template <typename T> ostream &operator<<(ostream &os, const vector<T> &v)
{
    os << "[ ";
    for (const auto &item : v)
        os << item << " ";
    os << "]\n";
    return os;
}

class Solution
{
  public:
    int jumpBFS(vector<int> &nums)
    {
        queue<pair<int, int>> q;
        int ans = INT32_MAX, targetIndex = nums.size() - 1;
        // times jumped, index
        q.push({0, 0});
        while (!q.empty())
        {
            pair<int, int> curr = q.front();
            q.pop();

            if (curr.second >= targetIndex)
            {
                ans = min(ans, curr.first);
                continue;
            }

            for (int i = curr.second + 1; i <= curr.second + nums[curr.second]; i++)
                q.push({curr.first + 1, i});
        }
        return ans;
    }
    int jump(vector<int> &nums)
    {
        // r represents the maximum index we can jump to.
        // l represents the minimum index we can start from.
        int l = 0, r = 0, n = nums.size();
        int ans = 0;
        while (r < n - 1)
        {
            int maxJump = 0;
            for (int i = l; i <= r; i++)
                // each time we check can we jump more? if we can update r
                if (i + nums[i] > maxJump)
                    maxJump = i + nums[i];
            l = r + 1;
            r = maxJump;
            ans++;
        }
        return ans;
    }
};

int main()
{
    auto s = new Solution();
    vector<int> nums = {2, 4, 1, 1, 1, 1};
    nums = {2, 1, 2, 1, 0};
    nums = {1, 2, 3, 4};
    cout << nums << s->jump(nums) << nl;
    return 0;
}