package item

import (
	"fmt"

	"github.com/graphql-go/graphql"
	"github.com/taskalla/api/pkg/db"
	"github.com/taskalla/api/pkg/models"
	"github.com/taskalla/api/pkg/paginate"
	"github.com/taskalla/api/pkg/tokenutils"
)

func UserItemsResolver(p graphql.ResolveParams, o paginate.ConnectionOptions) ([]paginate.Edge, *paginate.PageInfo, error) {
	t, err := tokenutils.ExtractToken(p)
	if err != nil {
		return nil, nil, err
	}

	var after *int
	if a, ok := o.After.(int); ok {
		after = &a
	}

	fmt.Println(after)

	filter_struct := ItemFilter{}
	if filter_map, ok := p.Args["filter"].(map[string]interface{}); ok {
		// There's a filter object
		if filter_done, ok := filter_map["done"].(bool); ok {
			filter_struct.Done = &filter_done
		}
	}

	/* rows, err := db.DB.Query(context.Background(), `
	 *   WITH items AS (
	 *     SELECT ROW_NUMBER() OVER (ORDER BY created_at DESC) AS cursor, id, item_description, user_id, done, created_at FROM items
	 *     WHERE user_id = $1
	 *     AND done = coalesce($2, done)
	 *   ) SELECT * FROM items
	 *   WHERE cursor > coalesce($3, 0)
	 *   LIMIT $4
	 * `, t.UserID, filter_struct.Done, after, o.First+1)
	 * if err != nil {
	 *   return nil, nil, err
	 * } */

	items := []models.Item{}
	result := db.DB.Where("user_id = ?", t.UserID).Find(&items)
	if result.Error != nil {
		return nil, nil, result.Error
	}

	all_edges := []paginate.Edge{}

	for _, item := range items {
		all_edges = append(all_edges, paginate.Edge{
			Cursor: 6,
			Node:   item,
		})
	}

	/* edges := all_edges
	 * if len(edges) != 0 && len(all_edges) > o.First {
	 *   edges = all_edges[:len(all_edges)-1]
	 * } */

	return all_edges, &paginate.PageInfo{HasNextPage: len(all_edges) > o.First}, nil
}
