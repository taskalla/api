package db

import (
	"context"
	"os"

	"github.com/Matt-Gleich/logoru"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/taskalla/api/pkg/logging"
)

var DB *pgxpool.Pool

func Connect() {
	conn, err := pgxpool.Connect(context.Background(), os.Getenv("DB"))

	if err != nil {
		logging.Critical("Error connecting to database: " + err.Error())
	} else {
		logoru.Success("Successful database connection!")
		DB = conn
	}
}
