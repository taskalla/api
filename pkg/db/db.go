package db

import (
	"fmt"
	"net/url"
	"os"
	"strings"

	"github.com/Matt-Gleich/logoru"
	"github.com/taskalla/api/pkg/logging"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	db_url, err := url.Parse(os.Getenv("DB"))
	if err != nil {
		logging.Critical("Invalid Postgres connection string in the DB environment variable")
	}

	password, _ := db_url.User.Password()

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		db_url.Hostname(),
		db_url.User.Username(),
		password,
		strings.TrimPrefix(db_url.Path, "/"),
		db_url.Port(),
	)

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: dsn,
	}), &gorm.Config{})
	if err != nil {
		logging.Critical("Error connecting to database: " + err.Error())
	} else {
		DB = db
		logoru.Success("Successful database connection!")
	}
}
