# Definition for a  binary tree node
class TreeNode:
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None


# Definition for singly-linked list.
class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None


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

    def listToArray(self, A):
        array = []
        while A is not None:
            array.append(A.val)
            A = A.next
        return array

    def sortedListToBST(self, A):
        array = self.listToArray(A)
        return self.sortedArrayToBSTLeftRight(array, 0, len(array) - 1)


if __name__ == "__main__":
    five = ListNode(5)

    four = ListNode(4)
    four.next = five

    three = ListNode(3)
    three.next = four

    two = ListNode(2)
    two.next = three

    one = ListNode(1)
    one.next = two

    s = Solution()
    s.sortedListToBST(one)