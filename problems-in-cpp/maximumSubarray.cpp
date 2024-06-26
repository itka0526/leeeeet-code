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
    int maxSubArray(vector<int> &nums)
    {
        int res = 0;
        int ans = INT32_MIN;
        for (int i = 0; i < nums.size(); i++)
        {
            if (res < 0)
                res = 0;
            res += nums[i];
            ans = max(ans, res);
        }
        return ans;
    }
};

int main()
{
    vector<int> nums = {2, -3, 4, -2, 2, 1, -1, 4};
    nums = {-2, 1, -3, 4, -1, 2, 1, -5, 4};
    auto s = new Solution();
    cout << "Solution: " << s->maxSubArray(nums) << nl;
    return 0;
}