package paginate

import (
	"errors"

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
	PageInfo PageInfo      `json:"pageInfo"`
}

type ConnectionOptions struct {
	First int    `json:"first"`
	After string `json:"after"`
}

type PaginatedResolveFunc func(graphql.ResolveParams, ConnectionOptions) ([]Edge, PageInfo, error)

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

func (obj ConnectionObj) ResolveFunc(f PaginatedResolveFunc) graphql.FieldResolveFn {
	return func(p graphql.ResolveParams) (interface{}, error) {
		opts := ConnectionOptions{}

		if first, ok := p.Args["first"].(int); ok {
			opts.First = first
		} else {
			return nil, errors.New("Please provide the `first` argument")
		}

		if after, ok := p.Args["after"].(string); ok {
			opts.After = after
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
