import networkx as nx

graph = nx.Graph()

for l in open("./in", "r"):
    l, r = l.split(":")
    for n in r.strip().split():
        graph.add_edge(l, n)

graph.remove_edges_from((nx.minimum_edge_cut(graph)))
a, b = nx.connected_components(graph)
print(len(a) * len(b))