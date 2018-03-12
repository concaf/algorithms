class StackNode:
    def __init__(self, data):
        self.data = data
        self.next = None

class MinStack:
    def __init__(self):
        self.root = None
        self.current_min = StackNode(-1)

    # @param x, an integer
    def push(self, x):
        if self.current_min.data == -1:
            self.current_min = StackNode(x)
        elif x < self.current_min.data:
            self.current_min = StackNode(data=x)

        current = self.root
        self.root = StackNode(data=x)
        self.root.next = current

    # @return nothing
    def pop(self):
        if self.is_empty():
            return

        popping = self.root
        self.root = self.root.next

        if self.current_min.data == popping.data:
            self.current_min = self.find_min()

    def find_min(self):
        current = self.root
        min_element = StackNode(-1)
        while current is not None:
            if min_element.data == -1:
                min_element = current
            elif current.data < min_element.data:
                min_element = current
            current = current.next
        return min_element

    # @return an integer
    def top(self):
        return -1 if self.root is None else self.root.data

    def is_empty(self):
        return True if self.root is None else False

    # @return an integer
    def getMin(self):
        return self.current_min.data

    def print_stack(self):
        current = self.root
        while True:
            if current is None:
                print("None")
                break
            print(str(current.data)+"->")
            current = current.next


if __name__ == '__main__':
    stack = MinStack()
    stack.push(19)
    stack.push(10)
    stack.push(9)
    stack.print_stack()
    print(stack.getMin())
    stack.push(8)
    print(stack.getMin())
    # print(stack.previous_min.data)
