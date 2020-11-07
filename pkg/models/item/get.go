package item

import (
	"context"

	"github.com/taskalla/api/pkg/db"
)

func GetUserItems(user string, count, page int, filter ItemFilter) ([]*Item, error) {
	rows, err := db.DB.Query(context.Background(), "SELECT id, item_description, user_id, done FROM items WHERE user_id = $1 AND done = coalesce($2, done) ORDER BY id ASC LIMIT $3 OFFSET $4", user, filter.Done, count, (page-1)*count)
	if err != nil {
		return nil, err
	}

	items := []*Item{}

	for rows.Next() {
		item := &Item{}
		err := rows.Scan(&item.ID, &item.Description, &item.UserID, &item.Done)
		if err != nil {
			return nil, err
		}

		items = append(items, item)
	}

	return items, nil
}

func GetItemCountOnPage(user string, count, page int) (int, error) {
	row := db.DB.QueryRow(context.Background(), "SELECT COUNT(*) AS count FROM (SELECT id FROM items WHERE user_id = $1 ORDER BY id ASC LIMIT $2 OFFSET $3) AS items", user, count, (page-1)*count)

	var returned_count int

	err := row.Scan(&returned_count)
	if err != nil {
		return 0, err
	}

	return returned_count, nil
}

func GetTotalItemCount(user string) (int, error) {
	row := db.DB.QueryRow(context.Background(), "SELECT COUNT(*) AS count FROM (SELECT id FROM items WHERE user_id = $1) AS items", user)

	var returned_count int

	err := row.Scan(&returned_count)
	if err != nil {
		return 0, err
	}

	return returned_count, nil
}
