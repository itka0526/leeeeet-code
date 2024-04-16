#include <bits/stdc++.h>

using namespace std;

// Naive approach
// int main()
// {
//     int64_t n, a, b, w, h;
//     cin >> n >> a >> b >> w >> h;
//     // Allocatable possible space for each module
//     int64_t e = w * h / n;
//     int64_t t = 0;
//     while (1)
//     {
//         if ((a + 2) * (b + 2) >= e)
//             break;
//         a += 2;
//         b += 2;
//         t++;
//     };
//     cout << t << endl;
//     return 0;
// }

int main()
{
    int64_t n, a, b, w, h;
    cin >> n >> a >> b >> w >> h;
    int64_t left = 0;
    int64_t right = 1000000000000000000LL + 1;
    while (left + 1 < right) {
        int64_t mid, d;
        mid = d = (left + right) / 2;
        int64_t cnt0 = (w / (a + 2 * d)) * (h / (b + 2 * d));
        int64_t cnt1 = (w / (b + 2 * d)) * (h / (a + 2 * d));
        if (cnt0 < n && cnt1 < n)
            right = mid;
        else
            left = mid;
    }
    cout << left << endl;
    return 0;
}