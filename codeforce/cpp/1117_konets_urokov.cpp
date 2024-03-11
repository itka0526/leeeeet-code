// #include <bits/stdc++.h>

// using namespace std;

// int main()
// {
//     int n;
//     cin >> n;
//     int p = 0;
//     // 5 + 15 + 5 + 15 + 5 + 15 ...
//     int k = 1;
//     while (k < n)
//     {
//         p += k % 2 == 1 ? 5 : 15;
//         k++;
//     }
//     p += n * 45;
//     cout << 9 + p / 60 << " " << p % 60;
//     return 0;
// }

#include <bits/stdc++.h>

using namespace std;

int main()
{
    int n;
    cin >> n;
    int s = 0;
    // Add the hours
    s += 45 * n;
    // Then we calculate assuming all the breaks are short and add them up by
    s += (n - 1) * 5;
    // Then we correct our the times by adding the missing long breaks differences
    s += (n - 1) / 2 * (15 - 5);
    cout << 9 + s / 60 << " " << s % 60;
    return 0;
}