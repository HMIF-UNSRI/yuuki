package mariadb

import (
	"database/sql"
	"fmt"
	"strconv"
	"time"
	"yuuki/pkg/config"
	"yuuki/pkg/exception"
)

func GetConnection(configuration config.Configuration) *sql.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		configuration.Get("MYSQL_USERNAME"),
		configuration.Get("MYSQL_PASSWORD"),
		configuration.Get("MYSQL_HOST"),
		configuration.Get("MYSQL_PORT"),
		configuration.Get("MYSQL_DATABASE"),
	)

	db, err := sql.Open("mysql", dsn)
	exception.PanicIfErr(err)

	poolMin, err := strconv.Atoi(configuration.Get("MYSQL_POOL_MIN"))
	exception.PanicIfErr(err)
	poolMax, err := strconv.Atoi(configuration.Get("MYSQL_POOL_MAX"))
	exception.PanicIfErr(err)

	db.SetMaxIdleConns(poolMin)
	db.SetMaxOpenConns(poolMax)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	return db
}
