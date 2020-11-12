package paginate

import "github.com/graphql-go/graphql"

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
	First *int `json:"first"`
}

type PaginatedResolveFunc func(graphql.ResolveParams, ConnectionOptions) (interface{}, PageInfo, error)

type ConnectionObj struct {
	Object *graphql.Object
	Args   graphql.FieldConfigArgument
}

func (obj ConnectionObj) ResolveFunc(f PaginatedResolveFunc) graphql.FieldResolveFn {
	return func(p graphql.ResolveParams) (interface{}, error) {
		opts := ConnectionOptions{}

		if first, ok := p.Args["first"].(int); ok {
			opts.First = &first
		}
		_, _, err := f(p, opts)
		if err != nil {
			return nil, err
		}
		return nil, nil
	}
}

var PageInfoObj = graphql.NewObject(graphql.ObjectConfig{
	Name: "PageInfo",
})

func NewConnectionObject(name string, wraps *graphql.Object) *ConnectionObj {
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
			},
		}),
		Args: graphql.FieldConfigArgument{
			"first": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
		},
	}
}
