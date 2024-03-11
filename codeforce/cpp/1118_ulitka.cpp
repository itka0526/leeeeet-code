// #include <bits/stdc++.h>

// using namespace std;

// int main()
// {
//     int h, a, b;
//     cin >> h >> a >> b;
//     // 3 - 2    1
//     // + 3 - 2  1 + 3 - 2 = 2
//     // + 3 - 2  2 + 3 - 2 = 3
//     // + 3 - 2  3 + 3 - 2 = 4
//     // + 3 - 2  4 + 3 - 2 = 5
//     // + 3 - 2  5 + 3 - 2 = 8 - 2 <- this day
//     int curr = 0;
//     int count = 0;
//     while (h > curr)
//     {
//         curr += a;
//         count += 1;
//         if (h <= curr)
//         {
//             break;
//         }
//         curr -= b;
//     }
//     cout << count;
//     return 0;
// }

#include <bits/stdc++.h>

using namespace std;

int main()
{
    int h, a, b, n;
    n = 1;
    cin >> h >> a >> b;
    if (a < h)
    {
        // Remaining distance divided by velocity
        n += (h - a) / (a - b);
        // If there is a remaining distance left we increment it by one
        if ((h - a) % (a - b) > 0)
            n++;
    }
    cout << n;
    return 0;
}