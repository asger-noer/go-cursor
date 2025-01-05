# Cursor [![Continuous Integration](https://github.com/asger-noer/go-cursor/actions/workflows/ci.yml/badge.svg)](https://github.com/asger-noer/go-cursor/actions/workflows/ci.yml) [![Go Report Card](https://goreportcard.com/badge/github.com/asger-noer/go-cursor)](https://goreportcard.com/report/github.com/asger-noer/go-cursor) [![Go Reference](https://pkg.go.dev/badge/github.com/asger-noer/go-cursor.svg)](https://pkg.go.dev/github.com/asger-noer/go-cursor)

Cursor is a generic implementation of the [Relay Cursor Connections Specification][relay_graphql_connection]. It is a simple library that can be used to paginate any list of items.

## Installation

Simply run the following command to get the package and start using it:

```bash
go get github.com/asger-noer/go-cursor
```

## Usage

You can create a new connection by providing a list of items and a cursor function that can be used to get the cursor for each item. You can also provide a list of arguments to the connection to specify the pagination.

```go
args := []cursor.Argument{
	cursor.First(10),
	cursor.After("MQ=="),
}

// Create a new connection with the list of users
cur, err := cursor.New(users, userCursor, args...)
if err != nil {
	// Handle error
}

// pageinfo contains information about the current page
pageinfo := cur.PageInfo()

// Edges contains the list of items in the current page
edges := cur.Edges() {

```

<!-- External links -->

[relay_graphql_connection]: https://relay.dev/graphql/connections.htm

### GQLGen

> [!TIP]
> You can find a complete example of how to use Cursor with gqlgen in the [examples/gqlgen](examples/gqlgen) directory.

If you are using [gqlgen](https://gqlgen.com/), you can use the `Cursor` type to generate the cursor for each item. You'll need to define the concrete implementation of the `Connection` and `Edge` types with the model being paginated. This is done by adding creating a type and binding it in the model section of the `gqlgen.yml` file:

_types.go:_
```go
package types

import (
	"github.com/asger-noer/go-cursor"
	"github.com/asger-noer/go-cursor/examples/gqlgen/graph/model"
)

type (
	TodoConnection = cursor.Connection[model.Todo]
	TodoEdge       = cursor.Edge[model.Todo]
)
```

_gqlgen.yml:_
```yaml
models:
  PageInfo:
    model:
      - github.com/asger-noer/go-cursor.PageInfo
  TodoConnection:
    model:
      - examples/gqlgen/graph/types.TodoConnection
  TodoEdge:
    model:
      - examples/gqlgen/graph/types.TodoEdge
```
