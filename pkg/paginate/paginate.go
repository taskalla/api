package paginate

import (
	"errors"
	"reflect"

	"github.com/graphql-go/graphql"
)

type Edge struct {
	Cursor string `json:"cursor"`
	Node   interface{}
}

type PageInfo struct {
	HasNextPage bool `json:"hasNextPage"`
}

type Connection struct {
	Nodes    []interface{} `json:"nodes"`
	Edges    []Edge        `json:"edges"`
	PageInfo *PageInfo     `json:"pageInfo"`
}

type ConnectionOptions struct {
	First int     `json:"first"`
	After *string `json:"after"`
}

type PaginatedResolveFunc func(graphql.ResolveParams, ConnectionOptions) ([]Edge, *PageInfo, error)
type SimplePaginatedResolveFunc func(graphql.ResolveParams) (interface{}, error)

type ConnectionObj struct {
	Object *graphql.Object
	Args   graphql.FieldConfigArgument
}

func edgesToNodes(edges []Edge) []interface{} {
	nodes := []interface{}{}

	for _, e := range edges {
		nodes = append(nodes, e.Node)
	}

	return nodes
}

func GetConnectionOptions(p graphql.ResolveParams) (ConnectionOptions, error) {
	opts := ConnectionOptions{}

	if first, ok := p.Args["first"].(int); ok {
		if first < 1 {
			return ConnectionOptions{}, errors.New("The `first` argument must be positive, and non-zero")

		}
		opts.First = first
	} else {
		return ConnectionOptions{}, errors.New("Please provide the `first` argument")
	}

	if after, ok := p.Args["after"].(string); ok {
		opts.After = &after
	}

	return opts, nil
}

func (obj ConnectionObj) ResolveFunc(f PaginatedResolveFunc) graphql.FieldResolveFn {
	return func(p graphql.ResolveParams) (interface{}, error) {
		opts, err := GetConnectionOptions(p)
		if err != nil {
			return nil, err
		}

		edges, pageInfo, err := f(p, opts)
		if err != nil {
			return nil, err
		}
		return Connection{
			Edges:    edges,
			PageInfo: pageInfo,
			Nodes:    edgesToNodes(edges),
		}, nil
	}
}

func (obj ConnectionObj) SimpleResolveFunc(f SimplePaginatedResolveFunc) graphql.FieldResolveFn {
	return func(p graphql.ResolveParams) (interface{}, error) {
		opts, err := GetConnectionOptions(p)
		if err != nil {
			return nil, err
		}

		// Deserialize provided cursor, if any
		cursor_index := -1
		if opts.After != nil {
			cursor_index, err = DeserializeCursor(*opts.After)
			if err != nil {
				return nil, err
			}
		}

		nodes_interface, err := f(p)
		if err != nil {
			return nil, err
		}

		nodes_value := reflect.ValueOf(nodes_interface)

		if nodes_value.Kind() != reflect.Slice {
			return nil, errors.New("Value is not a slice")
		}

		nodes := make([]Edge, nodes_value.Len())

		for i := 0; i < nodes_value.Len(); i++ {
			nodes[i] = Edge{
				Node:   nodes_value.Index(i).Interface(),
				Cursor: SerializeCursor(i),
			}
		}

		node_slice := []Edge{}

		pagination_range := [2]int{cursor_index + 1, opts.First + cursor_index + 1}

		page_info := PageInfo{
			HasNextPage: true,
		}

		if pagination_range[0] < len(nodes) {
			// It's valid!

			if pagination_range[1] >= len(nodes) {
				pagination_range[1] = len(nodes)
				page_info.HasNextPage = false
			}

			node_slice = nodes[pagination_range[0]:pagination_range[1]]
		} else {
			page_info.HasNextPage = false
		}

		return Connection{
			Nodes:    edgesToNodes(node_slice),
			Edges:    node_slice,
			PageInfo: &page_info,
		}, nil
	}
}

var PageInfoObj = graphql.NewObject(graphql.ObjectConfig{
	Name: "PageInfo",
	Fields: graphql.Fields{
		"hasNextPage": &graphql.Field{
			Type: graphql.NewNonNull(graphql.Boolean),
		},
	},
})

func newEdgeObject(name string, wraps *graphql.Object) *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: name,
		Fields: graphql.Fields{
			"cursor": &graphql.Field{
				Type: graphql.NewNonNull(graphql.String),
			},
			"node": &graphql.Field{
				Type: wraps,
			},
		},
	})
}

func NewConnectionObject(name string, wraps *graphql.Object, additionalArgs graphql.FieldConfigArgument) *ConnectionObj {
	args := graphql.FieldConfigArgument{
		"first": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
		"after": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	}

	for a, b := range additionalArgs {
		args[a] = b
	}

	return &ConnectionObj{
		Object: graphql.NewObject(graphql.ObjectConfig{
			Name: name,
			Fields: graphql.Fields{
				"nodes": &graphql.Field{
					Type: graphql.NewList(wraps),
				},
				"pageInfo": &graphql.Field{
					Type: PageInfoObj,
				},
				"edges": &graphql.Field{
					Type: graphql.NewList(newEdgeObject(name+"Edge", wraps)),
				},
			},
		}),
		Args: args,
	}
}
