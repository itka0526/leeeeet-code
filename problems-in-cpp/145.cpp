struct TreeNode
{
    int val;
    TreeNode *left;
    TreeNode *right;
    TreeNode() : val(0), left(nullptr), right(nullptr)
    {
    }
    TreeNode(int x) : val(x), left(nullptr), right(nullptr)
    {
    }
    TreeNode(int x, TreeNode *left, TreeNode *right) : val(x), left(left), right(right)
    {
    }
};

#include <bits/stdc++.h>

using namespace std;

class Solution
{
  public:
    vector<int> ans;
    void dfs(TreeNode *node)
    {
        if (node == NULL)
            return;
        dfs(node->left);
        dfs(node->right);
        ans.push_back(node->val);
    }
    vector<int> postorderTraversal(TreeNode *root)
    {
        dfs(root);
        return ans;
    }
};