#include <bits/stdc++.h>

using namespace std;

// int main()
// {
//     int n;
//     cin >> n;
//     vector<vector<int>> x(n, vector<int>(n));
//     for (int i = 0; i < n; i++)
//         for (int j = 0; j < n; j++)
//             cin >> x[i][j];

//     set<int> vis;
//     int roads = 0;

//     for (int i = 0; i < n; i++)
//     {
//         for (int j = 0; j < n; j++)
//         {
//             if (i == j || vis.find((i + 1) * (j + 1)) != vis.end())
//                 continue;

//             vector<int> a = x[i];
//             vector<int> b = x[j];

//             for (int g = 0; g < n; g++)
//                 if (a[g] == 1 && a[g] == b[g])
//                     roads++;

//             vis.insert((i + 1) * (j + 1));
//         }
//     }

//     cout << roads << endl;
//     return 0;
// }

#include <bits/stdc++.h>

using namespace std;

int main()
{
    int n;
    cin >> n;
    int mn = n * n, sum = 0;
    while (mn--)
    {
        int k;
        cin >> k;
        sum += k;
    }
    cout << sum / 2 << endl;
    return 0;
}