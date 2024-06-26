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
    bool canJumpBFS(vector<int> &nums)
    {
        queue<pair<int, int>> q;
        // index, jumpCount
        q.push({0, nums[0]});

        int targetIndex = nums.size() - 1;

        while (!q.empty())
        {
            pair<int, int> curr = q.front();
            q.pop();

            if (curr.first >= targetIndex)
                return true;

            for (int i = curr.first + 1; i <= curr.first + curr.second; i++)
                q.push({i, nums[i]});
        }
        return false;
    }
    bool canJump(vector<int> &nums)
    {
        int targetIndex = nums.size() - 1;
        // From the end of the array can we reach the start
        for (int i = nums.size() - 2; i >= 0; i--)
            // Update the new targetIndex, the main logic is that
            // if we can reach the prior element then we can totally reach the next
            if (i + nums[i] >= targetIndex)
                targetIndex = i;
        return targetIndex == 0;
    }
};

int main()
{
    auto s = new Solution();
    vector<int> nums = {1, 2, 0, 1, 0};
    nums = {1, 2, 1, 0, 1};
    cout << s->canJump(nums) << nl;
    return 0;
}