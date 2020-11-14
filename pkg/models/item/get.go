package item

import (
	"context"

	"github.com/graphql-go/graphql"
	"github.com/taskalla/api/pkg/db"
	"github.com/taskalla/api/pkg/paginate"
)

func GetUserItems(user string, count, page int, filter ItemFilter) ([]*Item, error) {
	rows, err := db.DB.Query(context.Background(), "SELECT id, item_description, user_id, done, created_at FROM items WHERE user_id = $1 AND done = coalesce($2, done) ORDER BY id ASC LIMIT $3 OFFSET $4", user, filter.Done, count, (page-1)*count)
	if err != nil {
		return nil, err
	}

	items := []*Item{}

	for rows.Next() {
		item := &Item{}
		err := rows.Scan(&item.ID, &item.Description, &item.UserID, &item.Done, &item.CreatedAt)
		if err != nil {
			return nil, err
		}

		items = append(items, item)
	}

	return items, nil
}

func UserItemsResolver(p graphql.ResolveParams, o paginate.ConnectionOptions) ([]paginate.Edge, paginate.PageInfo, error) {
	return nil, paginate.PageInfo{}, nil
}
