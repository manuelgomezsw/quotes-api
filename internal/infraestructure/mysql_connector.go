package infraestructure

import (
	"database/sql"
	"log"
	"sync"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var (
	onceConnectToDB sync.Once
	ClientDB        *sql.DB
)

func init() {
	onceConnectToDB.Do(
		func() {
			mySqlClientDB, err := connectToDB()
			if err != nil {
				log.Fatalf("Something happend %s", err)
			}
			ClientDB = mySqlClientDB
		},
	)
}

func connectToDB() (*sql.DB, error) {
	mySqlClientDB, err := sql.Open("mysql", "")
	if err != nil {
		return nil, err
	}

	mySqlClientDB.SetConnMaxLifetime(time.Minute * 3)
	mySqlClientDB.SetMaxOpenConns(10)
	mySqlClientDB.SetMaxIdleConns(10)

	err = mySqlClientDB.Ping()
	if err != nil {
		return nil, err
	}

	return mySqlClientDB, nil
}
