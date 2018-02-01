#
# https://leetcode.com/problems/binary-tree-right-side-view/description/
#
class TreeLinkNode:
    def __init__(self, x, l, r):
        self.val = x
        self.left = l
        self.right = r

class Solution:
    def rightSideView(self, root):
        if root is None:
            return []
        levels = {}
        q = [(root, 0)]

        while len(q) > 0:
            pair = q.pop(0)
            node, l = pair[0], pair[1]
            levels[l] = node

            if node.left is not None:
                q.append((node.left, l+1))
            if node.right is not None:
                q.append((node.right, l+1))

        return [levels[k].val for k in levels.keys()]


t = TreeLinkNode(1,
    TreeLinkNode(2,
        TreeLinkNode(4, None, None),
        TreeLinkNode(5, None, None)),
    TreeLinkNode(3, None, None)
)

print(Solution().rightSideView(t))
