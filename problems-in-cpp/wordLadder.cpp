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
    bool oneCharDiff(string w1, string w2)
    {
        // Lengths are always the same
        int n = w1.size();
        int cnt = 0;
        for (int i = 0; i < n; i++)
        {
            if (w1[i] != w2[i])
            {
                cnt += 1;
            }
            if (cnt > 1)
            {
                return false;
            }
        }
        return true;
    }

    int ladderLength(string beginWord, string endWord, vector<string> &wordList)
    {
        // If the endWord is not the list then its impossible to reach it!
        if (!count(wordList.begin(), wordList.end(), endWord))
        {
            return 0;
        }
        map<string, set<string>> m;
        // Add the beginWord in the list
        wordList.push_back(beginWord);
        // Create the graph
        for (const string &w1 : wordList)
        {
            for (const string &w2 : wordList)
            {
                if (w1 != w2 && oneCharDiff(w1, w2))
                {
                    m[w1].insert(w2);
                }
            }
        }
        // Debugging...
        for (pair<string, set<string>> p : m)
        {
            cout << p.first << ": ";
            for (const string &word : p.second)
            {
                cout << word << " ";
            }
            cout << '\n';
        }
        // BFS:
        // Keep the track of the travelled distance
        queue<pair<string, int>> q;
        q.push({beginWord, 1});
        int ans = 0;
        set<string> vis;
        while (!q.empty())
        {
            pair<string, int> curr = q.front();
            q.pop();
            // Our graph is bi-directional
            // If the node was visited just ignore it, so we don't create infinite loops
            if (vis.count(curr.first))
            {
                continue;
            }
            // If the word matches with endWord then we exit
            if (curr.first == endWord)
            {
                ans = curr.second;
                break;
            }
            // Add the node to visited
            vis.insert(curr.first);
            // Add the connected nodes to the queue
            for (string conn : m[curr.first])
            {
                q.push({conn, curr.second + 1});
            }
        }
        return ans;
    }
};

int main()
{
    auto s = new Solution();
    vector<string> wl = {"bat", "bag", "sag", "dag", "dot"};
    vector<string> wl1 = {"bat", "bag", "sat", "dag", "dot"};
    vector<string> wl2 = {"hot", "dot", "dog", "lot", "log", "cog"};
    auto ans = s->ladderLength("cat", "sag", wl);
    cout << "\n";
    auto ans1 = s->ladderLength("cat", "sag", wl1);
    cout << "\n";
    auto ans2 = s->ladderLength("hit", "cog", wl2);
    cout << "\n";

    cout << "ANSWER: " << ans << '\n';
    cout << "ANSWER 1: " << ans1 << '\n';
    cout << "ANSWER 2: " << ans2 << '\n';
    return 0;
}