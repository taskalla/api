package paginate

import (
	"errors"
	"reflect"

	"github.com/graphql-go/graphql"
	"github.com/taskalla/api/pkg/logging"
)

type Edge struct {
	Cursor interface{} `json:"cursor"`
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
	First int         `json:"first"`
	After interface{} `json:"after"`
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
		opts.First = first
	} else {
		return ConnectionOptions{}, errors.New("Please provide the `first` argument")
	}

	if after, ok := p.Args["after"]; ok {
		opts.After = after
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

		logging.Info(opts)

		nodes_interface, err := f(p)
		if err != nil {
			return nil, err
		}

		nodes_value := reflect.ValueOf(nodes_interface)

		if nodes_value.Kind() != reflect.Slice {
			return nil, errors.New("Value is not a slice")
		}

		nodes := make([]interface{}, nodes_value.Len())

		for i := 0; i < nodes_value.Len(); i++ {
			nodes[i] = nodes_value.Index(i).Interface()
		}

		return Connection{
			Nodes:    nodes,
			Edges:    []Edge{},
			PageInfo: &PageInfo{},
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

func newEdgeObject(name string, wraps *graphql.Object, cursorType graphql.Output) *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: name,
		Fields: graphql.Fields{
			"cursor": &graphql.Field{
				Type: graphql.NewNonNull(cursorType),
			},
			"node": &graphql.Field{
				Type: wraps,
			},
		},
	})
}

func NewConnectionObject(name string, wraps *graphql.Object, cursorType graphql.Output, additionalArgs graphql.FieldConfigArgument) *ConnectionObj {
	args := graphql.FieldConfigArgument{
		"first": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
		"after": &graphql.ArgumentConfig{
			Type: cursorType,
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
					Type: graphql.NewList(newEdgeObject(name+"Edge", wraps, cursorType)),
				},
			},
		}),
		Args: args,
	}
}
