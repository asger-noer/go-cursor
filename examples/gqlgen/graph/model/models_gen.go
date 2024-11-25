// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type PageInfo struct {
	EndCursor       *string `json:"endCursor,omitempty"`
	StartCursor     *string `json:"startCursor,omitempty"`
	HasNextPage     bool    `json:"hasNextPage"`
	HasPreviousPage bool    `json:"hasPreviousPage"`
}

type Query struct {
}

type Todo struct {
	ID   string `json:"id"`
	Text string `json:"text"`
	Done bool   `json:"done"`
	User *User  `json:"user"`
}

type TodoConnection struct {
	Edges    []*TodoEdge `json:"Edges,omitempty"`
	PageInfo *PageInfo   `json:"PageInfo,omitempty"`
}

type TodoEdge struct {
	Cursor string `json:"cursor"`
	Node   *Todo  `json:"node"`
}

type TodoInput struct {
	First  *int    `json:"first,omitempty"`
	After  *string `json:"after,omitempty"`
	Last   *int    `json:"last,omitempty"`
	Before *string `json:"before,omitempty"`
}

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}