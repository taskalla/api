package item

import (
	"context"

	"github.com/taskalla/api/pkg/db"
)

func GetUserItems(user string, count, page int) ([]*Item, error) {
	rows, err := db.DB.Query(context.Background(), "SELECT id, title, item_description, user_id FROM items WHERE user_id = $1 ORDER BY id ASC LIMIT $2 OFFSET $3", user, count, page*count)
	if err != nil {
		return nil, err
	}

	items := []*Item{}

	for rows.Next() {
		item := &Item{}
		err := rows.Scan(&item.ID, &item.Title, &item.Description, &item.UserID)
		if err != nil {
			return nil, err
		}

		items = append(items, item)
	}

	return items, nil
}
