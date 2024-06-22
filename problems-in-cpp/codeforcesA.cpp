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
    os << " ]\n";
    return os;
}

template <typename T> ostream &operator<<(ostream &os, const vector<T> &v)
{
    os << "[ ";
    for (const auto &item : v)
        os << item << " ";
    os << " ]\n";
    return os;
}

void creatingWords()
{
    string a, b;
    cin >> a >> b;
    swap(a[0], b[0]);
    cout << a << " " << b << nl;
}

void myFirstSortingProblem()
{
    int x, y;
    cin >> x >> y;

    if (x > y)
    {
        cout << y << " " << x;
    }
    else
    {
        cout << x << " " << y;
    }
    cout << nl;
}

void vladAndTheBestOfFive()
{
    int a = 0, b = 0;
    for (int i = 0; i < 5; i++)
    {
        char t;
        cin >> t;
        if (t == 'A')
            a++;
        else if (t == 'B')
            b++;
    }
    cout << (a > b ? 'A' : 'B') << nl;
}

void shortSort()
{
    string s;
    cin >> s;
    if (s[0] == 'a' && s[1] == 'b' && s[2] == 'c')
        YES;
    else if ((s[0] == 'a' && s[1] == 'c' && s[2] == 'b') || (s[0] == 'b' && s[1] == 'a' && s[2] == 'c') ||
             (s[0] == 'c' && s[1] == 'b' && s[2] == 'a'))
        YES;
    else
        NO;

    // int cnt = 0;
    // string letters = "abc";
    // for (int i = 0; i < 3; i++)
    //     cnt += (s[i] != letters[i]);
    // cnt <= 2 ? YES : NO;
}

void toMyCritics()
{
    vector<int> a(3);
    int sum = 0;
    for (int &num : a)
    {
        int t;
        cin >> t;
        sum += t;
        num = t;
    };
    for (int i = 0; i < 3; i++)
    {
        if (sum - a[i] >= 10)
        {
            YES;
            return;
        }
    }
    NO;
    // int a, b, c; cin >> a >> b >> c;
    // (a + b + c - min({a, b, c})) >= 10 ? YES : NO;
}

void loveStory()
{
    string s;
    cin >> s;
    string cmp_s = "codeforces";
    int ans = 0;
    for (int i = 0; i < s.size(); i++)
        if (s[i] != cmp_s[i])
            ++ans;
    cout << ans << nl;
}

void plusOrMinus()
{
    int a, b, c;
    cin >> a >> b >> c;
    cout << (a + b == c ? '+' : '-') << nl;
}

void codeforcesChecking()
{
    set<char> c = {'c', 'o', 'd', 'e', 'f', 'o', 'r', 'c', 'e', 's'};
    char ch;
    cin >> ch;
    c.count(ch) ? YES : NO;
}

void followingDirections()
{
    int n;
    string s;
    pair<int, int> pos = {0, 0};
    cin >> n;
    cin >> s;
    for (int i = 0; i < n; i++)
    {
        if (s[i] == 'L')
            --pos.first;
        else if (s[i] == 'R')
            ++pos.first;
        else if (s[i] == 'D')
            --pos.second;
        else if (s[i] == 'U')
            ++pos.second;
        if (pos.first == 1 && pos.second == 1)
        {
            YES;
            return;
        };
    }
    NO;
}

int main()
{
    int t;
    cin >> t;
    while (t--)
    {
        // creatingWords();
        // myFirstSortingProblem();
        // vladAndTheBestOfFive();
        // shortSort();
        // toMyCritics();
        // loveStory();
        // plusOrMinus();
        // codeforcesChecking();
        followingDirections();
    }
    return 0;
}