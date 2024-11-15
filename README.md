# Cursor [![Continuous Integration](https://github.com/asger-noer/cursor/actions/workflows/ci.yml/badge.svg)](https://github.com/asger-noer/cursor/actions/workflows/ci.yml) [![Go Report Card](https://goreportcard.com/badge/github.com/asger-noer/cursor)](https://goreportcard.com/report/github.com/asger-noer/cursor) [![Go Reference](https://pkg.go.dev/badge/github.com/asger-noer/cursor.svg)](https://pkg.go.dev/github.com/asger-noer/cursor)

Cursor is a generic implementation of the [Relay Cursor Connections Specification][relay_graphql_connection]. It is a simple library that can be used to paginate any list of items.

## Installation

Simply run the following command to get the package and start using it:

```bash
go get github.com/asger-noer/cursor
```

## Usage

You can create a new connection by providing a list of items and a cursor function that can be used to get the cursor for each item. You can also provide a list of arguments to the connection to specify the pagination.

```go
args := []cursor.Argument{
	cursor.First(10),
	cursor.After("MQ=="),
}

// Create a new connection with the list of users
connection, err := cursor.NewConnection(users, userCursor, args...)
if err != nil {
	// Handle error
}

// pageinfo contains information about the current page
pageinfo := connection.PageInfo()

// Edges contains the list of items in the current page
edges := connection.Edges() {

```

<!-- External links -->

[relay_graphql_connection]: https://relay.dev/graphql/connections.htm
