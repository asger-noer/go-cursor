package cursor_test

import (
	"encoding/base64"
	"testing"

	"github.com/alecthomas/assert/v2"
	"github.com/asger-noer/go-cursor"
)

type UserNodeResolver struct {
	ID       string
	UserName string
}

func TestNewConnectionFromEdges(t *testing.T) {
	type want struct {
		nodes           []UserNodeResolver
		startCursor     *string
		endCursor       *string
		hasPreviousPage bool
		hasNextPage     bool
		err             error
	}

	type test struct {
		name string
		args []cursor.Argument
		data []UserNodeResolver
		want want
	}

	cursorFn := func(t UserNodeResolver) string {
		return base64.StdEncoding.EncodeToString([]byte(t.ID))
	}

	nodes := []UserNodeResolver{
		{ID: "1", UserName: "John Doe"},
		{ID: "2", UserName: "Jane Doe"},
		{ID: "3", UserName: "Alice Doe"},
	}

	tests := []test{
		{
			name: "connection without any arguments",
			args: []cursor.Argument{}, // no options
			data: nodes,
			want: want{
				nodes:           nodes,
				startCursor:     typeToPtr(cursorFn(nodes[0])),
				endCursor:       typeToPtr(cursorFn(nodes[2])),
				hasPreviousPage: false,
				hasNextPage:     false,
			},
		},
		{
			name: "connection with after argument",
			args: []cursor.Argument{
				cursor.After(typeToPtr(cursorFn(nodes[0]))),
			},
			data: nodes,
			want: want{
				nodes: []UserNodeResolver{
					nodes[1], // Jane Doe
					nodes[2], // Alice Doe
				},
				startCursor:     typeToPtr(cursorFn(nodes[1])),
				endCursor:       typeToPtr(cursorFn(nodes[2])),
				hasPreviousPage: false, // There is not efficient way to check if there is a previous page
				hasNextPage:     false,
			},
		},
		{
			name: "connection with before argument",
			args: []cursor.Argument{
				cursor.Before(typeToPtr(cursorFn(nodes[2]))),
			},
			data: nodes,
			want: want{
				nodes: []UserNodeResolver{
					nodes[0], // John Doe
					nodes[1], // Jane Doe
				},
				startCursor:     typeToPtr(cursorFn(nodes[0])),
				endCursor:       typeToPtr(cursorFn(nodes[1])),
				hasPreviousPage: false,
				hasNextPage:     false, // There is not efficient way to check if there is a next page
			},
		},
		{
			name: "connection with first argument",
			args: []cursor.Argument{
				cursor.First(typeToPtr(1)),
			},
			data: nodes,
			want: want{
				nodes: []UserNodeResolver{
					nodes[0], // John Doe
				},
				startCursor:     typeToPtr(cursorFn(nodes[0])),
				endCursor:       typeToPtr(cursorFn(nodes[0])),
				hasPreviousPage: false,
				hasNextPage:     true,
			},
		},
		{
			name: "connection with last argument",
			args: []cursor.Argument{
				cursor.Last(typeToPtr(1)),
			},
			data: nodes,
			want: want{
				nodes: []UserNodeResolver{
					nodes[2], // Alice Doe
				},
				startCursor:     typeToPtr(cursorFn(nodes[2])),
				endCursor:       typeToPtr(cursorFn(nodes[2])),
				hasPreviousPage: true,
				hasNextPage:     false,
			},
		},
		{
			name: "connection with after and first argument",
			args: []cursor.Argument{
				cursor.After(typeToPtr(cursorFn(nodes[0]))),
				cursor.First(typeToPtr(1)),
			},
			data: nodes,
			want: want{
				nodes: []UserNodeResolver{
					nodes[1], // Jane Doe
				},
				startCursor:     typeToPtr(cursorFn(nodes[1])),
				endCursor:       typeToPtr(cursorFn(nodes[1])),
				hasPreviousPage: false, // There is not efficient way to check if there is a previous page
				hasNextPage:     true,
			},
		},
		{
			name: "connection with before and last argument",
			args: []cursor.Argument{
				cursor.Before(typeToPtr(cursorFn(nodes[2]))),
				cursor.Last(typeToPtr(1)),
			},
			data: nodes,
			want: want{
				nodes: []UserNodeResolver{
					nodes[1], // Jane Doe
				},
				startCursor:     typeToPtr(cursorFn(nodes[1])),
				endCursor:       typeToPtr(cursorFn(nodes[1])),
				hasPreviousPage: true,
				hasNextPage:     false, // There is not efficient way to check if there is a next page
			},
		},
		{
			name: "connection with none existing after argument",
			args: []cursor.Argument{
				cursor.After(typeToPtr("some_invalid_cursor")),
			},
			data: nodes,
			want: want{
				nodes:           nil,
				startCursor:     nil,
				endCursor:       nil,
				hasPreviousPage: false,
				hasNextPage:     false,
			},
		},
		{
			name: "connection with none existing before argument",
			args: []cursor.Argument{
				cursor.Before(typeToPtr("some_invalid_cursor")),
			},
			data: nodes,
			want: want{
				nodes:           nil,
				startCursor:     nil,
				endCursor:       nil,
				hasPreviousPage: false,
				hasNextPage:     false,
			},
		},
		{
			name: "connection with first and last argument",
			args: []cursor.Argument{
				cursor.First(typeToPtr(1)),
				cursor.Last(typeToPtr(1)),
			},
			data: nodes,
			want: want{
				nodes:           nil,
				startCursor:     nil,
				endCursor:       nil,
				hasPreviousPage: false,
				hasNextPage:     false,
				err:             cursor.ErrConflictingArguments,
			},
		},
		{
			name: "connection with negative first argument",
			args: []cursor.Argument{
				cursor.First(typeToPtr(-1)),
			},
			data: nodes,
			want: want{
				nodes:           nil,
				startCursor:     nil,
				endCursor:       nil,
				hasPreviousPage: false,
				hasNextPage:     false,
				err:             cursor.ErrInvalidFirst,
			},
		},
		{
			name: "connection with negative last argument",
			args: []cursor.Argument{
				cursor.Last(typeToPtr(-1)),
			},
			data: nodes,
			want: want{
				nodes:           nil,
				startCursor:     nil,
				endCursor:       nil,
				hasPreviousPage: false,
				hasNextPage:     false,
				err:             cursor.ErrInvalidLast,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			conn, err := cursor.New(tt.data, cursorFn, tt.args...)

			if tt.want.err == nil {
				assert.NoError(t, err)
			} else {
				assert.Error(t, err)
				assert.EqualError(t, err, tt.want.err.Error())
			}

			pageinfo := conn.PageInfo()
			assert.Equal(t, tt.want.startCursor, pageinfo.StartCursor(), "start cursor")
			assert.Equal(t, tt.want.endCursor, pageinfo.EndCursor(), "end cursor")
			assert.Equal(t, tt.want.hasPreviousPage, pageinfo.HasPreviousPage(), "has previous page")
			assert.Equal(t, tt.want.hasNextPage, pageinfo.HasNextPage(), "has next page")

			for i, edge := range conn.Edges() {
				assert.Equal(t, tt.want.nodes[i], edge.Node())
			}
		})
	}
}

func typeToPtr[T any](t T) *T {
	return &t
}
