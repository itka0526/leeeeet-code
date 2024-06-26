#include <bits/stdc++.h>
#define nl '\n'
#define ll long long
#define YES cout << "YES" << nl
#define NO cout << "NO" << nl
#define vi vector<int>
#define vll vector<long long>

using namespace std;

template <typename T> ostream &operator<<(ostream &os, pair<T, T> p)
{
    os << "(" << p.first << ", " << p.second << ")";
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

template <typename T, typename Sequence> ostream &operator<<(ostream &os, queue<T, Sequence> &q)
{
    vector<T> v(q.size());
    int i = 0;
    while (!q.empty())
    {
        v[i++] = q.front();
        q.pop();
    }
    os << "Q: " << v;
    for (auto e : v)
        q.push(e);
    return os;
}

template <typename T, typename Container, typename Compare>
ostream &operator<<(ostream &os, priority_queue<T, Container, Compare> &pq)
{
    vector<T> v(pq.size());
    int i = 0;
    while (!pq.empty())
    {
        v[i++] = pq.top();
        pq.pop();
    }
    os << "PQ: " << v;
    for (auto e : v)
        pq.push(e);
    return os;
}

template <typename T, typename Compare, typename Allocator>
ostream &operator<<(ostream &os, set<T, Compare, Allocator> &s)
{
    os << "S: { ";
    for (const auto &item : s)
        os << item << " ";
    os << "}\n";
    return os;
}

void watermelon()
{
    int w;
    cin >> w;
    !(w % 2) && w > 2 ? YES : NO;
}

void wayTooLongWords()
{
    string word;
    cin >> word;
    word.size() <= 10 ? cout << word : cout << word[0] << to_string(word.size() - 2) << word[word.size() - 1];
    cout << nl;
}

void team()
{
    int n;
    cin >> n;
    int ans = 0;
    while (n--)
    {
        int a, b, c;
        cin >> a >> b >> c;
        ans += (a + b + c >= 2);
    }
    cout << ans << nl;
}

void bitPlusPlus()
{
    int x = 0, n;
    cin >> n;
    while (n--)
    {
        string ops;
        cin >> ops;
        if (ops == "X++" || ops == "++X")
            x++;
        else if (ops == "--X" || ops == "X--")
            --x;
    }
    cout << x << nl;
}

void nextRound()
{
    int k, n;
    cin >> n >> k;
    vector<int> players(n);
    for (int &pl : players)
        cin >> pl;
    // k-th player to 0 based index
    k--;
    auto ans = count_if(players.begin(), players.end(), [&](int score) { return score >= players[k] && score > 0; });
    cout << ans << nl;
}

void dominoPiling()
{
    int n, m;
    cin >> n >> m;
    cout << n * m / 2 << nl;
}

void beautifulMatrix()
{
    for (int i = 0; i < 5; i++)
    {
        for (int j = 0; j < 5; j++)
        {
            int t;
            cin >> t;
            if (t)
            {
                cout << abs(2 - i) + abs(2 - j);
                return;
            }
        }
    }
}

void petyaAndStrings()
{
    string s1, s2;
    cin >> s1 >> s2;
    int n = s1.size();
    for (int i = 0; i < n; i++)
    {
        s1[i] = tolower(s1[i]);
        s2[i] = tolower(s2[i]);
    }
    if (s1 > s2)
    {
        cout << 1;
    }
    else if (s1 < s2)
    {
        cout << -1;
    }
    else
    {
        cout << 0;
    }
}

void helpfulMaths()
{
    string original;
    cin >> original;
    priority_queue<int, vector<int>, greater<int>> pq;
    for (int i = 0; i < original.size(); i++)
    {
        if (original[i] == '+')
            continue;
        pq.push(original[i] - '0');
    }
    while (!pq.empty())
    {
        cout << pq.top();
        pq.pop();
        if (pq.size() >= 1)
            cout << '+';
    }
}

void wordCapitalization()
{
    string s;
    cin >> s;
    s[0] = toupper(s[0]);
    cout << s;
}

void boyOrAGirl()
{
    string s;
    set<char> uc;
    cin >> s;
    for (const char &ch : s)
        uc.insert(ch);
    string ans = !(uc.size() % 2) ? "CHAT WITH HER!" : "IGNORE HIM!";
    cout << ans << nl;
}

int main()
{
    // team();
    // watermelon();
    // bitPlusPlus();
    // nextRound();
    // dominoPiling();
    // petyaAndStrings();
    // helpfulMaths();
    // bitPlusPlus();
    // wordCapitalization();
    boyOrAGirl();
    return 0;
}