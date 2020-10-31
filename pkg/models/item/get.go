package item

import (
	"context"

	"github.com/taskalla/api/pkg/db"
)

func GetUserItems(user string) ([]*Item, error) {
	rows, err := db.DB.Query(context.Background(), "SELECT id, title, item_description, user_id FROM items WHERE user_id = $1", user)
	if err != nil {
		return nil, err
	}

	items := []*Item{}

	for rows.Next() {
		item := &Item{}
		rows.Scan(&item.ID, &item.Title, &item.Description, &item.UserID)

		items = append(items, item)
	}

	return items, nil
}
