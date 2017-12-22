class Graph:
    def __init__(self, graph):
        self.graph = graph

    def neighbors(self, vertex):
        return self.graph[vertex]

    def edge_exists(self, vertex1, vertex2):
        return vertex2 in self.graph[vertex1]

    def vertices(self):
        return self.graph.keys()


def append_to_stack_and_mark_as_black(stack, state, vertex):
    # append the stack with current vertex
    stack.append(vertex)

    # mark the state of the current vertex as black
    state[vertex] = 1


def next_white_vertex(graph, vertex, state):
    for neighbor in graph.neighbors(vertex):
        # if neighbor is black, continue to beginning of the loop
        if state[neighbor] == 1:
            continue

        # since the neighbor is white, set it as the current_vertex
        return neighbor
    return False


def pop_and_set_current_vertex_to_one_before(stack):
    stack.pop()
    return stack[-1]


def dfs(start_vertex, graph):
    visited_stack = []
    current_state = {}
    predecessor_graph = []

    # initializing current state of all vertices to be 0
    # 0 = white
    # 1 = black
    for vertex in graph.vertices():
        current_state[vertex] = 0

    # initialize current vertex to the given start vertex
    current_vertex = start_vertex

    append_to_stack_and_mark_as_black(visited_stack,
                                      current_state,
                                      current_vertex)

    while True:
        previous_vertex = current_vertex
        current_vertex = next_white_vertex(graph, current_vertex, current_state)
        if current_vertex:
            predecessor_graph.append("{} -> {}".format(current_vertex,
                                                       previous_vertex))

        # if no white neighbor is available, set it to the one before
        if not current_vertex:
            try:
                current_vertex = \
                    pop_and_set_current_vertex_to_one_before(visited_stack)
            except IndexError:
                break

        if current_vertex != visited_stack[-1]:
            append_to_stack_and_mark_as_black(visited_stack,
                                              current_state,
                                              current_vertex)

        if len(graph.neighbors(current_vertex)) == 0:
            try:
                current_vertex = \
                    pop_and_set_current_vertex_to_one_before(visited_stack)
            except IndexError:
                break

        if len(visited_stack) == 0:
            break

    for predecessor in predecessor_graph:
        print(predecessor)

if __name__ == "__main__":
    graph_representation = {
        "A": ["B", "D"],
        "B": ["C"],
        "C": ["B", "D", "E"],
        "D": ["A", "C", "E"],
        "E": ["C", "D", "F", "G"],
        "F": ["E"],
        "G": ["E"],
        "X": ["Y"],
        "Y": ["X"]
    }
    g = Graph(graph_representation)
    dfs(g.vertices()[0], g)
