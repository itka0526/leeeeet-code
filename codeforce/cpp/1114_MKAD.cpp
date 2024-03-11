#include <bits/stdc++.h>

using namespace std;

int main()
{
    int v, t;
    cin >> v >> t;
    int road_length = 109;
    int traveled = v * t;
    int cycles = traveled % road_length;
    cout << (cycles + road_length) % road_length + 1;
    return 0;
}