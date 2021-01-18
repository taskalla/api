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

	filter := map[string]interface{}{}
	if filter_map, ok := p.Args["filter"].(map[string]interface{}); ok {
		// There's a filter object

		filter = filter_map
	}

	filter["user_id"] = t.UserID

	items := []models.Item{}
	result := db.DB.Where(filter).Find(&items)
	if result.Error != nil {
		return nil, result.Error
	}

	return items, nil
}
