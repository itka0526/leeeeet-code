#include <bits/stdc++.h>
#define nl '\n'
#define ll long long
#define YES cout << "YES" << nl
#define NO cout << "NO" << nl
#define vi vector<int>
#define vll vector<long long>
using namespace std;

// Definition for a Node.
class Node
{
  public:
    int val;
    vector<Node *> neighbors;
    Node()
    {
        val = 0;
        neighbors = vector<Node *>();
    }
    Node(int _val)
    {
        val = _val;
        neighbors = vector<Node *>();
    }
    Node(int _val, vector<Node *> _neighbors)
    {
        val = _val;
        neighbors = _neighbors;
    }
};

class Solution
{
  public:
    Node *dfs(Node *node, unordered_map<Node *, Node *> &m)
    {
        if (node == nullptr)
        {
            return nullptr;
        }
        // If exists just return. Full circle? Bi-direcitonal?
        if (m.count(node))
        {
            return m[node];
        }
        Node *clone = new Node(node->val);
        m[node] = clone;
        for (Node *neighbor : node->neighbors)
        {
            clone->neighbors.push_back(dfs(neighbor, m));
        }
        return clone;
    }
    Node *cloneGraphDFS(Node *node)
    {
        unordered_map<Node *, Node *> oldToNew;
        return dfs(node, oldToNew);
    }

    Node *cloneGraphFAIL(Node *node)
    {
        unordered_map<int, set<int>> m;
        queue<Node *> q;
        if (node == nullptr)
            return nullptr;
        q.push(node);

        while (!q.empty())
        {
            Node *curr = q.front();
            q.pop();
            if (m.count(curr->val))
                continue;
            for (Node *neighbor : curr->neighbors)
            {
                m[curr->val].insert(neighbor->val);
                q.push(neighbor);
            }
        }
        // Build?
        // Nodes have been generated?
        unordered_map<int, Node *> gen;
        for (pair<int, set<int>> entry : m)
        {
            Node *n = new Node(entry.first);
            gen[n->val] = n;
        };

        // Connect the nodes?
        for (pair<int, Node *> n : gen)
        {
            for (int nei : m[n.first])
            {
                Node *nn = gen[nei];
                n.second->neighbors.push_back(nn);
                nn->neighbors.push_back(n.second);
            }
        }
        return gen[node->val];
    }

    Node *cloneGraph(Node *node)
    {
        if (node == nullptr)
            return nullptr;
        unordered_map<int, Node *> gen;
        // If the node hasn't been generated create the node attach the neighbors?
        // If it has been generated connect it?
        queue<Node *> q;
        q.push(node);
        while (!q.empty())
        {
            const Node *curr = q.front();
            q.pop();
            if (gen.count(curr->val))
            {
                continue;
            }
            Node *clone = new Node(curr->val);
            gen[clone->val] = clone;
            for (Node *nei : curr->neighbors)
            {
                q.push(nei);
                if (gen.count(nei->val))
                {
                    // The node was already generated
                    clone->neighbors.push_back(gen[nei->val]);
                    gen[nei->val]->neighbors.push_back(clone);
                }
            }
        }
        return gen[node->val];
    }
};

int main()
{
    auto sol = new Solution();
    return 0;
}