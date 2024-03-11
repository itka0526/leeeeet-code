#!/usr/bin/env python3

import sys


def build_graph(arr):
    graph = {}
    i = 0
    while i < len(arr):
        node = arr[i]
        connections = []
        i += 1
        while i < len(arr) and arr[i] != node:
            connections.append(arr[i])
            i += 1
        graph[node] = connections
    return graph


n = int(sys.stdin.readline().strip())
lang = sys.stdin.readline().strip().split()
p = list(map(int, sys.stdin.readline().strip().split()))

graph = build_graph(p)
for node, connections in graph.items():
    print(f"Node {node} is connected to {connections}")
