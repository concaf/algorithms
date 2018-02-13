# Given a linked list, return the node where the cycle begins. If there is no cycle, return null.
#
# Try solving it using constant additional space.
#
# Example :
#
# Input :
#
#                   ______
#                  |     |
#                  \/    |
#         1 -> 2 -> 3 -> 4
#
# Return the node corresponding to node 3.

# Definition for singly-linked list.
class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None

class Solution:
    # @param A : head node of linked list
    # @return the first node in the cycle in the linked list
    def detectCycle(self, A):
        if A is None:
            return None

        if A.next is None:
            return None

        if A.next == A:
            return A

        current = A.next
        previous = A

        empty_next = ListNode("empty")

        while True:
            print "current is: " + str(current.val)

            if current.next is None:
                return None

            previous.next = empty_next
            print str(previous.val) + " now points to empty"

            if current.next.next == empty_next:
                return current.next

            previous = current
            current = current.next


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

    # loop
     five.next = three

     s = Solution()
     node = s.detectCycle(one)
     print "loop is at: " + str(node.val)