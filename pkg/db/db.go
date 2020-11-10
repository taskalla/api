package db

import (
	"context"
	"os"
	"time"

	"github.com/Matt-Gleich/logoru"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/taskalla/api/pkg/logging"
)

var DB *pgxpool.Pool

func Connect() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	conn, err := pgxpool.Connect(ctx, os.Getenv("DB"))

	if err != nil {
		logging.Critical("Error connecting to database: " + err.Error())
	} else {
		logoru.Success("Successful database connection!")
		DB = conn
	}
}
