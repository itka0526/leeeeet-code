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
    vector<vector<int>> kClosest(vector<vector<int>> &points, int k)
    {
        auto cmpPoints = [](const vector<int> &a, const vector<int> &b) {
            return a[0] * a[0] + a[1] * a[1] > b[0] * b[0] + b[1] * b[1];
        };
        priority_queue<vector<int>, vector<vector<int>>, decltype(cmpPoints)> pq(cmpPoints);
        vector<vector<int>> ans;
        for (const vector<int> &point : points)
            pq.push(point);
        for (int i = 0; i < k; i++)
        {
            ans.push_back(pq.top());
            pq.pop();
        }
        return ans;
    }
};
