#include <bits/stdc++.h>
#define nl '\n'
#define ll long long
#define YES cout << "YES" << nl
#define NO cout << "NO" << nl
#define vi vector<int>
#define vll vector<long long>
using namespace std;

class Solution
{
  public:
    int find(int node, vector<int> &parents)
    {
        int parent = parents[node];
        // Go up the graph till we find root node
        while (parent != parents[node])
        {
            // Shorten the path, so we create a tree-like structure, instead of linked list (optional)
            parents[parents[node]] = parents[node];
            // Go up
            parent = parents[node];
        }
        return parent;
    }
    bool Funion(int node1, int node2, vector<int> &parents, vector<int> &ranks)
    {
        // Find the roots of the unions
        int parent1 = find(parents[node1], parents), parent2 = find(parents[node2], parents);
        // If they match it means that they are already in one union,
        // therefore the current connection will be redundant and would create an unnecessary cycle
        if (parent1 == parent2)
        {
            return true;
        }
        // Bigger union swallows the smaller union
        if (ranks[parent1] > ranks[parent2])
        {
            parents[parent2] = parent1;
            ranks[parent1] += ranks[parent2];
        }
        else
        {
            parents[parent1] = parent2;
            ranks[parent2] += ranks[parent1];
        }
        return false;
    }
    vector<int> findRedundantConnection(const vector<vector<int>> &edges)
    {
        int n = edges.size() + 1;
        vector<int> parents(n);
        // Initially each node's rank is 1
        vector<int> ranks(n, 1);
        // And each node is a root itself
        for (int i = 0; i < edges.size(); i++)
        {
            parents[i] = i;
        }
        for (const vector<int> &edge : edges)
        {
            // Till we find a redundant connection loop through the graph
            if (Funion(edge[0], edge[1], parents, ranks))
            {
                return edge;
            };
        }
        return {};
    }
};

int main()
{
    return 0;
}