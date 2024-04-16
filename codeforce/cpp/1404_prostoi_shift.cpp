#include <bits/stdc++.h>

using namespace std;

int main()
{
    string s;
    cin >> s;
    transform(s.cbegin(), s.cend(), s.begin(), [](unsigned char c) { return (c - 'a' + 1) % 26 + 'a'; });
    cout << s;
    return 0;
}