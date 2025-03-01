package graph

// Traits represents a set of graph traits and types, such as directedness or acyclicness. These
// traits can be set when creating a graph by passing the corresponding functional options, for
// example:
//
//	g := graph.New(graph.IntHash, graph.Directed())
//
// This will set the IsDirected field to true.
type Traits struct {
	IsDirected bool
	IsAcyclic  bool
	IsWeighted bool
	IsRooted   bool
}

// Directed creates a directed graph. This has implications on graph traversal and the order of
// arguments of the Edge and AddEdge functions.
func Directed() func(*Traits) {
	return func(t *Traits) {
		t.IsDirected = true
	}
}

// Acyclic creates an acyclic graph, which won't allow adding directed edges resulting in a cycle.
func Acyclic() func(*Traits) {
	return func(t *Traits) {
		t.IsAcyclic = true
	}
}

// Weighted creates a weighted graph. To set weights, use the Edge and AddEdge functions.
func Weighted() func(*Traits) {
	return func(t *Traits) {
		t.IsWeighted = true
	}
}

// Rooted creates a rooted graph. This is particularly common for building tree data structures.
func Rooted() func(*Traits) {
	return func(t *Traits) {
		t.IsRooted = true
	}
}

// Tree is an alias for Acyclic and Rooted, since most trees in Computer Science are rooted trees.
func Tree() func(*Traits) {
	return func(t *Traits) {
		Acyclic()(t)
		Rooted()(t)
	}
}
