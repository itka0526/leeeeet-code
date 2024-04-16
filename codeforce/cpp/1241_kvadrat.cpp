#include <bits/stdc++.h>

using namespace std;

int main()
{
    int n;
    cin >> n;
    int root = sqrt(n);
    printf(n == root * root ? "YES" : "NO");
    return 0;
}