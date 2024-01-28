class TreeNode:
    def __init__(self):
        self.children = []
        self.descendants = 0

def build_tree(N, inheritances):
    nodes = [TreeNode() for _ in range(N)]

    for i, parent in enumerate(inheritances):
        if parent != -1:
            nodes[parent].children.append(i)

    return nodes

def update_descendants(node, nodes):
    node.descendants = 1  # Include itself
    for child in node.children:
        node.descendants += update_descendants(nodes[child], nodes)
    return node.descendants

def find_repository_with_max_inheritance(nodes):
    max_inheritance = -1
    repository_with_max_inheritance = -1

    for i, node in enumerate(nodes):
        if node.descendants > max_inheritance:
            max_inheritance = node.descendants
            repository_with_max_inheritance = i

    return repository_with_max_inheritance

# Read the input data
N = int(input())

inheritances = [int(input()) for _ in range(N)]

root = None
for i, parent in enumerate(inheritances):
    if i == parent:
        root = i
        break

# Build the binary tree
nodes = build_tree(N, inheritances)

update_descendants(nodes, nodes)

# Find the repository with the maximum inheritance
result = find_repository_with_max_inheritance(nodes)

# Print the output
print(result)
