package types

import (
	"examples/gqlgen/graph/model"

	"github.com/asger-noer/go-cursor"
)

type (
	TodoConnection = cursor.Connection[model.Todo]
	TodoEdge       = cursor.Edge[model.Todo]
)
