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
    priority_queue<int> pq;
    int lastStoneWeight(vector<int> &stones)
    {
        for (int stone : stones)
            pq.push(stone);

        while (pq.size() >= 2)
        {
            int s1 = pq.top();
            pq.pop();
            int s2 = pq.top();
            pq.pop();
            cout << "Current: " << s1 << " " << s2 << '\n';
            if (s1 == s2)
                continue;
            else
                pq.push(abs(s1 - s2));
        }
        return pq.size() == 0 ? 0 : pq.top();
    }
};
