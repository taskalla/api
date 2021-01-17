package item

import (
	"github.com/graphql-go/graphql"
	"github.com/taskalla/api/pkg/db"
	"github.com/taskalla/api/pkg/models"
	"github.com/taskalla/api/pkg/tokenutils"
)

func UserItemsResolver(p graphql.ResolveParams) (interface{}, error) {
	t, err := tokenutils.ExtractToken(p)
	if err != nil {
		return nil, err
	}

	/* filter_struct := ItemFilter{}
	 * if filter_map, ok := p.Args["filter"].(map[string]interface{}); ok {
	 *   // There's a filter object
	 *   if filter_done, ok := filter_map["done"].(bool); ok {
	 *     filter_struct.Done = &filter_done
	 *   }
	 * } */

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
		return nil, result.Error
	}

	return items, nil
}
