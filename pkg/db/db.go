package db

import (
	"github.com/Matt-Gleich/logoru"
	"github.com/jackc/pgx"
	"github.com/taskalla/api/pkg/logging"
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
		logging.Critical("Error connecting to database: " + err.Error())
	} else {
		logoru.Success("Successful database connection!")
	}
}
