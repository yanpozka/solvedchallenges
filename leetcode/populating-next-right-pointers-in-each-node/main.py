# https://leetcode.com/problems/populating-next-right-pointers-in-each-node/description/
# https://leetcode.com/problems/populating-next-right-pointers-in-each-node-ii/description/
#


class TreeLinkNode:
    def __init__(self, x, l, r):
        self.val = x
        self.left = l
        self.right = r
        self.next = None


class Solution:

    def connect(self, root):
        if root is None:
            return
        levels = {}
        q = [(root, 0)]

        while len(q) > 0:
            pair = q.pop(0)
            node, l = pair[0], pair[1]
            lis = levels.get(l, [])
            lis.append(node)
            levels[l] = lis

            if node.left is not None:
                q.append((node.left, l+1))
            if node.right is not None:
                q.append((node.right, l+1))

        for k in levels.keys():
            nodes = levels[k]
            print(nodes)
            ix, lim = 0, len(nodes)-1
            while ix < lim:
                nodes[ix].next = nodes[ix+1]
                ix += 1


t = TreeLinkNode(1,
                 TreeLinkNode(2,
                              TreeLinkNode(4, None, None),
                              TreeLinkNode(5, None, None)),
                 TreeLinkNode(3,
                              TreeLinkNode(6, None, None),
                              TreeLinkNode(7, None, None))
                 )

print(Solution().connect(t))
