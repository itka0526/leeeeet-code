#include <bits/stdc++.h>
#define nl '\n'
#define ll long long
#define YES cout << "YES" << nl
#define NO cout << "NO" << nl
#define vi vector<int>
#define vll vector<long long>
using namespace std;

int main()
{
    return 0;
}

class Solution
{
  public:
    bool hasDuplicate(vector<int> &nums)
    {
        set<int> s(nums.begin(), nums.end());
        return s.size() != nums.size();
    }

    bool isAnagram(string s, string t)
    {
        if (s.size() != t.size())
        {
            return false;
        }
        sort(s.begin(), s.end());
        sort(t.begin(), t.end());
        return s == t;
    }

    vector<int> twoSumON2(vector<int> &nums, int target)
    {
        vector<int> ans;
        for (int i = 0; i < nums.size(); i++)
        {
            for (int j = 0; j < nums.size(); j++)
            {
                if (i != j && nums[i] + nums[j] == target)
                {
                    ans.push_back(i);
                }
            }
        }
        return ans;
    }

    vector<int> twoSum(vector<int> &nums, int target)
    {
        vector<int> ans;

        return ans;
    }
};
