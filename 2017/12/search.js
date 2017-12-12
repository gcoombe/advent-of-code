class Graph {
    constructor() {
        this.edges = new Map();
    }

    addEdge(start, end) {
        this.edges.set(start, this.edges.get(start) ? this.edges.get(start).add(end) : new Set([end]));
        this.edges.set(end, this.edges.get(end) ? this.edges.get(end).add(start) : new Set([start]));
    }

    _searchNode(nodeId, visitedNodes) {
        visitedNodes.add(nodeId);
        this.edges.get(nodeId).forEach(adjacentNode => {
            if (!visitedNodes.has(adjacentNode)) {
                return this._searchNode(adjacentNode, visitedNodes);
            }
        });
    }

    removeNodes(nodes) {
        nodes.forEach(node => this.edges.delete(node));
    }

    isEmpty() {
        return this.edges.size === 0;
    }

    nodesInConnectedComponent(nodeId) {
        const nodes = new Set();
        this._searchNode(nodeId, nodes);
        return nodes;
    }

    getFirstNodeInGraph() {
        return this.edges.size && this.edges.keys().next().value;
    }
}

const addLine = (graph, line) => {
    const lineComponents = line.split(" <-> ");
    const start = lineComponents[0].trim();
    const ends = lineComponents[1].split(", ");

    ends.forEach(end => {
        graph.addEdge(start, end);
    });

    return graph;
};

module.exports.part1 = (input, startingNode) => {
    const lines = input.split("\n");

    const graph = new Graph();
    lines.forEach(line => addLine(graph, line));

    const connectedNodes = graph.nodesInConnectedComponent(startingNode);
    return connectedNodes.size;
};

module.exports.part2 = (input) => {
    const lines = input.split("\n");

    const graph = new Graph();
    lines.forEach(line => addLine(graph, line));

    let connectedComponentsCount = 0;
    while(!graph.isEmpty()) {
        connectedComponentsCount++;

        const connectedNodes = graph.nodesInConnectedComponent(graph.getFirstNodeInGraph());
        graph.removeNodes(connectedNodes);
    }
    return connectedComponentsCount;
};