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
    int leastInterval(vector<char> &tasks, int n)
    {
        unordered_map<char, int> m;
        priority_queue<int> pq;
        queue<pair<int, int>> q; // curr, idleTime
        int time = 0;

        for (char t : tasks)
            m[t]++;

        for (auto p : m)
            pq.push(p.second);

        while (!pq.empty() || !q.empty())
        {
            if (!q.empty() && time >= q.front().second)
            {
                pq.push(q.front().first);
                q.pop();
            }
            if (!pq.empty())
            {
                int cnt = pq.top() - 1;
                pq.pop();
                if (cnt > 0)
                    q.push({cnt, time + n + 1});
            }
            time++;
        }
        return time;
    }

    int leastIntervalON(vector<char> &tasks, int n)
    {
        int time = 0;
        unordered_map<char, int> m;
        for (char task : tasks)
            m[task]++;
        priority_queue<int> pq;
        for (pair<char, int> p : m)
            pq.push(p.second);
        // While its not idling keep adding the frequent element
        // if its idling we should queue other elements
        queue<pair<int, int>> q;
        while (!pq.empty())
        {
            if (!q.empty() && time >= q.front().second)
            {
                pq.push(q.front().first);
                q.pop();
            }
            if (!pq.empty())
            {
                int cnt = pq.top() - 1;
                pq.pop();
                if (cnt > 0)
                    q.push({cnt, time + n + 1});
            }
            time++;
        }
        return time;
    }
};

int main()
{
    vector<char> s1 = {'X', 'X', 'Y', 'Y'};
    vector<char> s2 = {'A', 'A', 'A', 'B', 'C'};
    vector<char> s3 = {'A', 'B', 'C', 'D', 'E'};
    auto t = new Solution();
    t->leastInterval(s1, 2);
    t->leastInterval(s2, 3);
    t->leastInterval(s3, 2);
    return 0;
}
