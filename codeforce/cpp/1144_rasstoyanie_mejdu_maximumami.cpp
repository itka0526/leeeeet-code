#include <bits/stdc++.h>

using namespace std;

int main()
{
    int a, b, c;
    vector<int> m;
    cin >> a >> b >> c;
    int i = 1;
    while (a && b && c)
    {
        if (b > a && b > c)
            m.push_back(i);
        a = b;
        b = c;
        cin >> c;
        i += 1;
    }
    if (m.size() < 2)
    {
        cout << 0 << endl;
    }
    else
    {
        int minimum = INT32_MAX;
        for (int g = 0; g < m.size(); g++)
        {
            for (int h = 0; h < m.size(); h++)
            {
                if (g == h)
                    continue;
                int s = abs(m[g] - m[h]);
                if (s < minimum)
                    minimum = s;
            }
        }
        cout << minimum << endl;
    }
    return 0;
}