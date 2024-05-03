#include <bits/stdc++.h>

using namespace std;

class NumArray {
   private:
    vector<int> prefixSum;

   public:
    NumArray(vector<int>& nums) {
        prefixSum.resize(nums.size());
        prefixSum[0] = nums[0];
        for (int i = 1; i < prefixSum.size(); i++) {
            prefixSum[i] = prefixSum[i - 1] + nums[i];
        }
    }

    int sumRange(int left, int right) {
        if (left == 0) {
            return prefixSum[right];
        }
        return prefixSum[right] - prefixSum[left - 1];
    }
};

int main() {
    vector<vector<int>> s = {{{-2, 0, 3, -5, 2, -1}}, {0, 2}, {2, 5}, {0, 5}};
    auto na = new NumArray(s[0]);
    for (int i = 1; i < s.size(); i++) {
        cout << na->sumRange(s[i][0], s[i][1]) << endl;
    }
    return 0;
}
