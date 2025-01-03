package cursor

import (
	"errors"
)

var (
	ErrConflictingArguments = errors.New("first and last cannot be used together")
	ErrInvalidFirst         = errors.New("first argument must be a non-negative integer")
	ErrInvalidLast          = errors.New("last argument must be a non-negative integer")
)

// Connection is a generic connection type that is used to paginate results.
type Connection[T any] struct {
	edges    Edges[T]
	pageInfo PageInfo
}

// New creates a new Pages object. The pages object is used to paginate
func New[T any](nodes []T, cursor func(T) string, args ...Argument) (Connection[T], error) {
	var err error
	var arguments arguments

	for _, arg := range args {
		if argErr := arg(&arguments); argErr != nil {
			err = errors.Join(err, argErr)
		}
	}

	if err != nil {
		// We want to parse all the arguments before returning the error.
		return Connection[T]{
			edges: Edges[T]{},
		}, err
	}

	edges := newEdges(nodes, cursor)
	trimmed := edges.applyCursor(arguments.after, arguments.before)

	trimmed, err = trimmed.trimEdges(arguments.first, arguments.last)
	if err != nil {
		return Connection[T]{
			edges: Edges[T]{},
		}, err
	}

	return Connection[T]{
		edges: trimmed,
		pageInfo: PageInfo{
			startCursor:     trimmed.startCursor(),
			endCursor:       trimmed.endCursor(),
			hasPreviousPage: trimmed.hasPreviousPage(edges, arguments),
			hasNextPage:     trimmed.hasNextPage(edges, arguments),
		},
	}, nil
}

// PageInfo returns the page info for the connection.
func (p *Connection[T]) PageInfo() PageInfo {
	return p.pageInfo
}

// Edges returns the edges for the connection.
func (p *Connection[T]) Edges() []Edge[T] {
	return p.edges
}

type PageInfo struct {
	startCursor     *string
	endCursor       *string
	hasPreviousPage bool
	hasNextPage     bool
}

// StartCursor returns the start cursor for the connection.
func (p *PageInfo) StartCursor() *string {
	return p.startCursor
}

// EndCursor returns the end cursor for the connection.
func (p *PageInfo) EndCursor() *string {
	return p.endCursor
}

// HasPreviousPage returns true if there is a previous page.
func (p *PageInfo) HasPreviousPage() bool {
	return p.hasPreviousPage
}

// HasNextPage returns true if there is a next page.
func (p *PageInfo) HasNextPage() bool {
	return p.hasNextPage
}

type Argument func(*arguments) error

type arguments struct {
	first  *int    // first limits the number of results returned from the after cursor.
	after  *string // after is the cursor to start from.
	last   *int    // last limits the number of results returned up to the the before cursor.
	before *string // before is the cursor to start from.
}

// Before sets the before option for the connection.
func Before(before *string) Argument {
	return func(p *arguments) error {
		p.before = before
		return nil
	}
}

// After sets the after option for the connection.
func After(after *string) Argument {
	return func(p *arguments) error {
		p.after = after
		return nil
	}
}

// First sets the first option for the connection.
func First(first *int) Argument {
	return func(p *arguments) error {
		if p.last != nil {
			return ErrConflictingArguments
		}

		p.first = first
		return nil
	}
}

// Last sets the last option for the connection.
func Last(last *int) Argument {
	return func(p *arguments) error {
		if p.first != nil {
			return ErrConflictingArguments
		}

		p.last = last
		return nil
	}
}
