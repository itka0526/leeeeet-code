#include <bits/stdc++.h>

struct TreeNode {
    int val;
    TreeNode *left;
    TreeNode *right;
    TreeNode() : val(0), left(nullptr), right(nullptr) {}
    TreeNode(int x) : val(x), left(nullptr), right(nullptr) {}
    TreeNode(int x, TreeNode *left, TreeNode *right) : val(x), left(left), right(right) {}
};

using namespace std;

// 94. Binary Tree Inorder Traversal

class Solution {
   public:
    vector<int> ans;

    void dfs(TreeNode *node) {
        if (node != nullptr) {
            dfs(node->left);
            ans.push_back(node->val);
            dfs(node->right);
        }
    };

    vector<int> inorderTraversal(TreeNode *root) {
        dfs(root);
        return ans;
    }
};