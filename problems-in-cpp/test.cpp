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
    priority_queue<int> pq;
    for (int i = 1; i < 10; i++)
        pq.push(i * i);

    while (!pq.empty())
    {
        std::cout << pq.top() << " ";
        pq.pop();
    }
    std::cout << '\n';
    return 0;
}