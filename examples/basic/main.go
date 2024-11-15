package main

import (
	"fmt"

	"github.com/asger-noer/cursor"
)

type Users struct {
	ID       string
	UserName string
}

func userCursor(user Users) string {
	return user.ID
}

var users []Users = []Users{
	{ID: "1", UserName: "Alice"},
	{ID: "2", UserName: "Bob"},
	{ID: "3", UserName: "Charlie"},
	{ID: "4", UserName: "David"},
	{ID: "5", UserName: "Eve"},
	// ...
}

func main() {
	var (
		first = 2
		after = "2"
	)

	args := []cursor.Argument{
		cursor.First(&first),
		cursor.After(&after),
	}

	// Create a new connection with the list of users
	connection, err := cursor.NewConnection(users, userCursor, args...)
	if err != nil {
		// Handle error
	}

	for _, user := range connection.Edges() {
		fmt.Println("cursor:", user.Cursor(), "\tnode:\t", user.Node())
		// prints:
		// cursor: 3 	node:	 {3 Charlie}
		// cursor: 4 	node:	 {4 David}
	}

	pageinfo := connection.PageInfo()
	fmt.Println(
		"startCursor:", *pageinfo.StartCursor(),
		"endCursor:", *pageinfo.EndCursor(),
		"hasPreviousPage:", pageinfo.HasPreviousPage(),
		"hasNextPage:", pageinfo.HasNextPage(),
	)
	// prints:
	// startCursor: 3 endCursor: 4 hasPreviousPage: false hasNextPage: true
}
