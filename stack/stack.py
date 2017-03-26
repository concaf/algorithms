class Stack(object):

    def __init__(self):
        self.items = []

    def __repr__(self):
        return str(self.items)

    def push(self, item):
        self.items.append(item)

    def pop(self):
        return self.items.pop()

    def peek(self):
        return self.items[-1]

    def is_empty(self):
        return len(self.items) == 0

    def size(self):
        return len(self.items)


def main():
    stack = Stack()
    stack.push("I")
    stack.push("Am")
    stack.push("Batman")
    print("This is the stack: {}".format(stack))
    print("The top item in the stack is {}".format(stack.peek()))
    print("Is the stack empty?: {}".format(stack.is_empty()))
    print("The stack holds {} items right now".format(stack.size()))
    print("Removing {} from the stack, and putting in Shinchan".
          format(stack.pop()))
    stack.push("Shinchan")
    print("The final stack is {}".format(stack))


if __name__ == '__main__':
    main()
