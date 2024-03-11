#include <bits/stdc++.h>

using namespace std;

// 5
// 3 3
// 6 1
// 10 2
// 10 3

int main()
{
    int a;
    cin >> a;
    int_fast64_t c = 0;
    while (a-- > 0)
    {
        // time and duration
        int t, d;
        cin >> t >> d;

        // If I am watching something
        if (t >= c)
            c = t + d;
        else
            c += d;
    }

    cout << c;
    return 0;
}
