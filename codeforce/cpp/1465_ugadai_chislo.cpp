#include <bits/stdc++.h>

using namespace std;

int main()
{
    char op;
    int64_t left = 1, right = 1e9;
    while (1)
    {
        int64_t curr = (left + right) / 2;

        cout << curr << endl;
        cin >> op;

        if (op == '>')
            // Unknown is higher than current
            left = curr + 1;

        else if (op == '<')
            // Unknown is lower than current
            right = curr - 1;

        else if (op == '=')
            // Found the number
            break;
    }
    return 0;
}
