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

// class KthLargest
// {
//   public:
//     int k;
//     vector<int> nums;
//     KthLargest(int k, const vector<int> &nums)
//     {
//         this->k = k;
//         this->nums = nums;
//     }

//     int add(int val)
//     {
//         nums.push_back(val);
//         sort(nums.begin(), nums.end(), greater<int>());
//         return nums[k - 1];
//     }
// };

class KthLargest
{
  public:
    int k;
    priority_queue<int, vector<int>, greater<int>> pq;

    KthLargest(int k, vector<int> &nums)
    {
        this->k = k;
        for (int n : nums)
        {
            // We must initially push it to the queue
            pq.push(n);
            // Afterwards PQ will reorder itself and if the size exceeds k we can just remove the last element
            // so we can use top() to retrieve k'th largest element.
            if (pq.size() > k)
                pq.pop();
        }
    }

    int add(int value)
    {
        pq.push(value);
        // Remove the unneccessary element
        if (pq.size() > k)
            pq.pop();
        return pq.top();
    }
};