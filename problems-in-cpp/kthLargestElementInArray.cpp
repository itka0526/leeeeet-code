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
    int findKthLargest(vector<int> &nums, int k)
    {
        priority_queue<int, vector<int>, greater<int>> pq;
        for (int num : nums)
        {
            pq.push(num);
            if (pq.size() > k)
                pq.pop();
        }
        return pq.top();
    }
};

int main()
{
    auto sol = new Solution();
    vector<int> s = {2, 3, 1, 5, 4};
    sol->findKthLargest(s, 2);
    return 0;
}
