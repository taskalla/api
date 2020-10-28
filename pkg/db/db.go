package db

import (
	"github.com/Matt-Gleich/logoru"
	"github.com/jackc/pgx"
)

func Connect() {
	_, err := pgx.Connect(pgx.ConnConfig{
		Host:     "db",
		Database: "taskalla",
		Port:     5432,
		User:     "postgres",
		Password: "password",
	})
	if err != nil {
		logoru.Critical(err)
	} else {
		logoru.Success("success!")
	}
}
