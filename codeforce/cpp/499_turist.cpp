#include <bits/stdc++.h>

using namespace std;

int main()
{
    int p, w;
    cin >> p >> w;
    vector<int> weight(3);
    vector<int> people(3);
    cin >> weight[0] >> people[0] >> weight[1] >> people[1] >> weight[2] >> people[2];

    // 5kg = 6 people
    // 5kg = 4 people
    // 6kg = 5 people

    // Naive approach
    // int ck = 0, cw = 0;
    // int i = 0;
    // while (ck < k && cw < w)
    // {
    //     ck += b[i];
    //     cw += a[i];
    //     i++;
    // }
    // ck == k &&cw == w ? cout << "YES" : cout << "NO";
    // return 0;
    // Two conditions must be met

    // Better version I guess?
    for (int i = 0; i < 3; i++)
    {
        if (weight[i] <= w && people[i] >= p)
        {
            cout << "YES";
            return 0;
        }
        for (int j = 0; j < 3; j++)
        {
            if (i == j)
                continue;
            if (weight[i] + weight[j] <= w && people[i] + people[j] >= p)
            {
                cout << "YES";
                return 0;
            }
            for (int k = 0; k < 3; k++)
            {
                if (i == k || j == k)
                    continue;
                if (weight[i] + weight[j] + weight[k] <= w && people[i] + people[j] + people[k] >= p)
                {
                    cout << "YES";
                    return 0;
                }
            }
        }
    }
    cout << "NO";
    return 0;
}