#include <bits/stdc++.h>

using namespace std;

int main()
{
    ios::sync_with_stdio(false);
    cin.tie(NULL);
    cout.tie(NULL);

    int a, b, c, m = 0;
    cin >> a >> b >> c;
    while (a && b && c)
    {
        if (b > a && b > c)
            m += 1;
        a = b;
        b = c;
        cin >> c;
    }
    cout << m;
    return 0;
}

// HOLY FUCKING B*TCH
// import sys
// k = 0
// nums = sys.stdin.readline().strip().split(' ')
// idx = nums.index('0')
// nums = list(map(int, nums[:idx]))
// if len(nums) < 3:
//     print(0)
// else:
//     a = nums[0]
//     b = nums[1]
//     for i in range(2, len(nums)):
//         c = nums[i]
//         if b > a and b > c:
//             k += 1
//         a = b
//         b = c
//     print(k)
