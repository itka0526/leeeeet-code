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

void solveA()
{
    vector<int> coords(3);
    for (int &x : coords)
        cin >> x;
    int ans = INT32_MAX;
    for (int a = 1; a <= 10; a++)
        ans = min(ans, abs(coords[0] - a) + abs(coords[1] - a) + abs(coords[2] - a));
    cout << ans << endl;
}

bool isLocalMaxima(int i, int j, const vector<vector<int>> &mat)
{
    int val = mat[i][j];
    if (i > 0 && mat[i - 1][j] > val)
        return false;
    if (i < mat.size() - 1 && mat[i + 1][j] > val)
        return false;
    if (j > 0 && mat[i][j - 1] > val)
        return false;
    if (j < mat[0].size() && mat[i][j + 1] > val)
        return false;
    return true;
}

int localMinuma(int i, int j, const vector<vector<int>> &mat)
{
    if (j > 0)
        return mat[i][j - 1];
    if (i > 0)
        return mat[i - 1][j];
    if (j < mat[0].size() - 1)
        return mat[i][j + 1];
    if (i < mat.size() - 1)
        return mat[i + 1][j];
    return mat[i][j] - 1;
}

void solveB()
{
    int n, m;
    cin >> n >> m;
    vector<vector<int>> mat(n, vector<int>(m));
    for (int i = 0; i < n; i++)
        for (int j = 0; j < m; j++)
            cin >> mat[i][j];

    for (int i = 0; i < n; i++)
    {
        for (int j = 0; j < m; j++)
        {
            if (isLocalMaxima(i, j, mat))
            {
                mat[i][j] = localMinuma(i, j, mat);
            }
        }
    }

    for (int i = 0; i < n; i++)
    {
        for (int j = 0; j < m; j++)
        {
            cout << mat[i][j] << " ";
        }
        cout << nl;
    }
}

void solveC()
{
    int n1, n2;
    cin >> n1 >> n2;
    string s1, s2;
    cin >> s1;
    queue<int> q;
    for (int i = 0, num; i < n2; i++)
    {
        cin >> num;
        q.push(num);
    }
    cin >> s2;
    int i = 0;
    while (!q.empty())
    {
        int ind = q.front() - 1;
        q.pop();
        s1[ind] = s2[i];
        i++;
    }
    cout << s1 << nl;
}

int main()
{
    int t;
    cin >> t;
    while (t--)
    {
        // solveA();
        // solveB();
        solveC();
    }
    return 0;
}