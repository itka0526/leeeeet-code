# Techniques to solve tree problems

### Depth first search

You need to go through both branches of the tree using recursion: recurse node.Left, node.Right

### Using queue

as the name suggests its, it takes advantage of queues all though this one is not simple as one might think

It traverses through the tree using generic for loop and if there are nodes left it adds them to the queue and till the queue reaches the absolute size of 0, meaning the tree has been fully traversed, and while traversing the tree it does work and pops the element off

### Swapping branches

This task is much more simpler it just requires basic knowledge of recursion and how storing variables works. And I think the code speaks for itself.

left := invertTree(root.Left)
right := invertTree(root.Right)

### Checking if trees are equal

This one quite intuitive too. It recurses over the nodes and just checks if the searching tree's value is same as the main big tree. And if big tree and searching tree are traversed all the way till nil that means there is a match inside the big tree. And to check if not just check if one of the trees values don't match or reached nil too soon then return false. And we do this using recursion.

### Find the common node

This one is requires bit of thinking but its simple too. It uses recursion and bit of knowledge of how binary trees work each branch spreads into two branches but the branches are not completely random they are based on comparison on the right we have higher values and on the left its lesser value on so on and so forth.

The code just speaks for itself :D

if p.Val > root.Val && q.Val > root.Val {
// we must dive deeper to the right of the tree because p and q have common ancestor on that side
return lowestCommonAncestor(root.Right, p, q)
}

if p.Val < root.Val && q.Val < root.Val {
// same goes here too !
return lowestCommonAncestor(root.Left, p, q)
}

return root
