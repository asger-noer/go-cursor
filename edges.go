package cursor

import (
	"slices"
)

// Edge is a generic type that is used to define the edges for a
// connection in at paginated result. The Edge interface is used to define the
// the cursor that is used to paginate the nodes and the node itself.
type Edge[T any] struct {
	node     T
	cursorFn func(T) string
}

func (e Edge[T]) Node() T {
	return e.node
}

func (e Edge[T]) Cursor() string {
	return e.cursorFn(e.node)
}

func newEdges[T any](nodes []T, cursorFunc func(T) string) Edges[T] {
	var edges []Edge[T]

	for _, node := range nodes {
		edges = append(edges, Edge[T]{
			node:     node,
			cursorFn: cursorFunc,
		})
	}

	return edges
}

type Edges[T any] []Edge[T]

func (e Edges[T]) Nodes() []T {
	var nodes []T

	for _, edge := range e {
		nodes = append(nodes, edge.Node())
	}

	return nodes
}

func (e Edges[T]) startCursor() *string {
	if len(e) == 0 {
		return nil
	}
	cursor := e[0].Cursor()
	return &cursor
}

func (e Edges[T]) endCursor() *string {
	if len(e) == 0 {
		return nil
	}
	cursor := e[len(e)-1].Cursor()
	return &cursor
}

func (e Edges[T]) hasPreviousPage(all Edges[T], args arguments) bool {
	if args.last != nil {
		edges := all.applyCursor(args.after, args.before)
		if len(edges) > int(*args.last) {
			return true
		}
	}

	// We currently don't have a way to determine if there are more items before
	// the first one hence we always return false.
	if args.after != nil {
		return false
	}

	return false
}

func (e Edges[T]) hasNextPage(all Edges[T], args arguments) bool {
	if args.first != nil {
		edges := all.applyCursor(args.after, args.before)

		if len(edges) > int(*args.first) {
			return true
		}
	}

	// We currently don't have a way to determine if there are more items after
	// the last one hence we always return false.
	if args.before != nil {
		return false
	}

	return false
}

// applyCursor applies the cursor to the edges and returns a new set of edges.
func (e Edges[T]) applyCursor(after, before *string) Edges[T] {
	var edges []Edge[T]

	switch {
	case after != nil:
		afterEdge := -1
		for i, edge := range e {
			if edge.Cursor() == *after {
				afterEdge = i
				break
			}
		}

		if afterEdge == -1 {
			return []Edge[T]{}
		}

		edges = e[afterEdge+1:]
	case before != nil:
		beforeEdge := -1
		for i, edge := range e {
			if edge.Cursor() == *before {
				beforeEdge = i
				break
			}
		}

		if beforeEdge == -1 {
			return []Edge[T]{}
		}

		edges = e[:beforeEdge]
	default:
		edges = e
	}

	return edges
}

// trimEdges is used to trim the edges based on the first or last arguments.
func (e Edges[T]) trimEdges(first, last *int) (Edges[T], error) {
	edges := slices.Clone(e)
	switch {
	case first != nil:
		if *first < 0 {
			return nil, ErrInvalidFirst
		}

		if len(edges) > *first {
			edges = edges[:*first]
		}
	case last != nil:
		if *last < 0 {
			return nil, ErrInvalidLast
		}

		if len(edges) > *last {
			edges = edges[len(edges)-*last:]
		}
	}

	return edges, nil
}
