#include <bits/stdc++.h>

using namespace std;

int main()
{
    int c, p, s = 1, ms = 1;
    cin >> p;
    while (p)
    {
        cin >> c;
        if (p == c)
            s += 1;
        else
        {
            if (s > ms)
                ms = s;
            s = 1;
        }
        p = c;
    }
    cout << ms << endl;
    return 0;
}
// main()
// {
//     int k = 1, m = 1, a, b;
//     cin >> b;
//     do
//     {
//         a = b;
//         cin >> b;
//         if (a == b)
//             k++;
//         else
//         {
//             if (k > m)
//                 m = k;
//             k = 1;
//         }
//     } while (b);
//     cout << m;
// }