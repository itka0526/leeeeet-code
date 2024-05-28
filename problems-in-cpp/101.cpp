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

class Solution
{
  public:
    bool isSymmetric(TreeNode *root)
    {
        if (root->left == NULL && root->right == NULL)
        {
            return true;
        }
        std::deque<TreeNode *> q;
        q.push_front(root->left);
        q.push_back(root->right);
        while (!q.empty())
        {
            auto f = q.front();
            auto b = q.back();
            q.pop_front();
            q.pop_back();

            // If both are null
            if (!f && !b)
                continue;
            // If one of them is null
            if (!f || !b)
                return false;
            // If the values do not match
            if (f->val != b->val)
                return false;

            q.push_front(f->left);
            q.push_back(b->right);

            q.push_front(f->right);
            q.push_back(b->left);
        }
        return true;
    }
};