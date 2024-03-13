#include <bits/stdc++.h>

using namespace std;

int main()
{
    int n;
    cin >> n;
    // Sequence of numbers
    // 1, 1, 2, 3, 5, 8, 13, 21, 34
    int t1 = 1, t2 = 1, t3 = 1;
    int i = 3;
    if (n == 2)
    {
        cout << 1 << '\n' << i << endl;
        return 0;
    }
    else
        while (t3 < n)
        {
            int k = t1;
            t1 = t2;
            t2 = k + t2;
            t3 = t1 + t2;
            i++;
        }
    if (t3 == n)
        cout << 1 << '\n' << i << endl;
    else
        cout << 0 << endl;
    return 0;
}

// #include <bits/stdc++.h>

// using namespace std;

// int main()
// {
//     int n;
//     cin >> n;
//     int f1, f2, f3;
//     f1 = f2 = f3 = 1;
//     int i = 2;
//     while (f3 < n)
//     {
//         f3 = f1 + f2;
//         f1 = f2;
//         f2 = f3;
//         i += 1;
//     }
//     if (f3 == n)
//         cout << 1 << '\n' << i << endl;
//     else
//         cout << 0 << endl;
//     return 0;
// }