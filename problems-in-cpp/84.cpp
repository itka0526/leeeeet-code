#include <bits/stdc++.h>
#define nl '\n'
#define ll long long
#define YES cout << "YES" << nl
#define NO cout << "NO" << nl
#define vi vector<int>
#define vll vector<long long>
using namespace std;

int main()
{
    return 0;
}

class Solution
{
  public:
    int largestRectangleArea(vector<int> &heights)
    {
        int ma = *max_element(heights.begin(), heights.end());
        for (int i = 0; i < heights.size(); i++)
        {
            int j = i;
            while (j < heights.size())
            {
                int y = *min_element(heights.begin() + i, heights.begin() + i + j);
                int x = j - i + 1;
                cout << x << " " << y << '\n';
                if (ma < x * y)
                {
                    ma = x * y;
                }
                j++;
            }
        }
        return ma;
    }
};
