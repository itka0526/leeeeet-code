#include <bits/stdc++.h>

using namespace std;

int main()
{
    int c;
    int sum = 0, n = 0;
    vector<int> nums;
    while (cin >> c && c != 0)
    {
        n += 1;
        sum += c;
        nums.push_back(c);
    }
    cout << fixed << setprecision(6);
    long double avg = sum * 1.0 / n;
    long double standartDeviationSum = 0.0;
    for (int i = 0; i < nums.size(); i++)
        standartDeviationSum += powl(nums[i] - avg, 2.0);
    cout << sqrtl(standartDeviationSum / (n - 1)) << endl;
    return 0;
}