# Definition for a  binary tree node
class TreeNode:
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None


class Solution:
    def sortedArrayToBSTLeftRight(self, A, left, right):
        if right < left:
            return None
        mid = (right + left) / 2
        rootNode = TreeNode(A[mid])
        rootNode.left = self.sortedArrayToBSTLeftRight(A, left, mid-1)
        rootNode.right = self.sortedArrayToBSTLeftRight(A, mid+1, right)
        return rootNode

    def sortedArrayToBST(self, A):
        return self.sortedArrayToBSTLeftRight(A, 0, len(A)-1)


if __name__ == "__main__":
    s = Solution()
    input = [1, 2, 3, 4, 5, 6, 7, 8, 9]
    print input
    s.sortedArrayToBST(input)
