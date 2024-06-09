#include <bits/stdc++.h>
#define nl '\n'
#define ll long long
#define YES cout << "YES" << nl
#define NO cout << "NO" << nl
#define vi vector<int>
#define vll vector<long long>
#define user unordered_map<int, vector<int>>
using namespace std;

class Twitter
{
  private:
    // user -> follows users
    unordered_map<int, set<int>> followsDb;
    // user -> tid, time
    unordered_map<int, vector<pair<int, int>>> tweetsDb;
    int time;

  public:
    Twitter()
    {
        time = 0;
    }

    void postTweet(int userId, int tweetId)
    {
        tweetsDb[userId].push_back({tweetId, time++});
    }

    vector<int> getNewsFeed(int userId)
    {
        // Compare function for priority queue
        auto cmp = [](const pair<int, int> &a, const pair<int, int> &b) { return a.second > b.second; };
        // Get the list of users the user follows
        // Self tweets should show up in feed
        followsDb[userId].insert(userId);
        set<int> follows = followsDb[userId];

        priority_queue<pair<int, int>, vector<pair<int, int>>, decltype(cmp)> pq(cmp);

        for (auto followee : follows)
        {
            // Then for each user it follows retrieve tweets
            vector<pair<int, int>> unorderedTweets = tweetsDb[followee];
            for (auto i : unorderedTweets)
                cout << i.first << " Time: " << i.second << "; \n";
            for (pair<int, int> tweet : unorderedTweets)
            {
                pq.push(tweet);
                if (pq.size() > 10)
                    pq.pop();
            }
        }

        vector<int> result;
        while (!pq.empty() && result.size() <= 10)
        {
            result.push_back(pq.top().first);
            pq.pop();
        }
        reverse(result.begin(), result.end());
        return result;
    }

    void follow(int followerId, int followeeId)
    {
        followsDb[followerId].insert(followeeId);
    }

    void unfollow(int followerId, int followeeId)
    {
        followsDb[followerId].erase(followeeId);
    }
};

int main()
{
    Twitter *twitter = new Twitter();
    twitter->postTweet(1, 1);
    twitter->postTweet(1, 2);
    twitter->postTweet(1, 3);
    twitter->postTweet(1, 4);
    twitter->postTweet(1, 5);
    twitter->postTweet(1, 6);
    twitter->postTweet(1, 7);
    twitter->postTweet(1, 8);
    twitter->postTweet(1, 9);
    twitter->postTweet(1, 10);
    twitter->postTweet(1, 11);
    vector<int> res = twitter->getNewsFeed(1);
    for (auto i : res)
        std::cout << i << ' ';
    cout << endl;
    twitter->follow(2, 1);
    twitter->getNewsFeed(2);
    vector<int> res2 = twitter->getNewsFeed(1);
    for (auto i : res2)
        std::cout << i << ' ';
    cout << endl;
    twitter->unfollow(2, 1);
    twitter->getNewsFeed(2);
    return 0;
}