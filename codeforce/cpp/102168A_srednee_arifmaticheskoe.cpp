#include <bits/stdc++.h>

using namespace std;

int main()
{
    ios::sync_with_stdio(false);
    cin.tie(NULL);
    cout.tie(NULL);
    cout << fixed << setprecision(9);
    int n;
    cin >> n;
    int64_t sum = 0, count = 0;
    while (n-- > 0)
    {
        char o;
        int x;
        cin >> o >> x;
        if (o == '+')
        {
            sum += x;
            count += 1;
        }
        else if (o == '-')
        {
            sum -= x;
            count -= 1;
        }
        cout << (count == 0 ? 0 : sum * 1.0 / count) << endl;
    }

    return 0;
}