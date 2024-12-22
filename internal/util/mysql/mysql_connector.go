package mysql

import (
	"cloud.google.com/go/cloudsqlconn"
	"context"
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net"
	"quotes-api/internal/infraestructure/client/secretmanager"
	"quotes-api/internal/util/constant"
	"sync"
)

var (
	ClientDB *sql.DB
	once     sync.Once
)

func init() {
	once.Do(func() {
		ClientDB = mustConnect()
	})
}

func mustConnect() *sql.DB {
	var err error
	var mysqlDB *sql.DB

	mysqlDB, err = connectWithConnector()
	if err != nil {
		log.Fatalf("connectConnector: unable to connect: %s", err)
	}

	if mysqlDB == nil {
		log.Fatal("Missing database connection type. Please define one of INSTANCE_HOST, INSTANCE_UNIX_SOCKET, or INSTANCE_CONNECTION_NAME")
	}

	return mysqlDB
}

func connectWithConnector() (*sql.DB, error) {
	dbUser, err := secretmanager.GetValue(constant.DbUser)
	if err != nil {
		return nil, err
	}

	dbPwd, err := secretmanager.GetValue(constant.DbPassword)
	if err != nil {
		return nil, err
	}

	dbName, err := secretmanager.GetValue(constant.DbName)
	if err != nil {
		return nil, err
	}

	instanceConnectionName, err := secretmanager.GetValue(constant.InstanceConnectionName)
	if err != nil {
		return nil, err
	}

	d, err := cloudsqlconn.NewDialer(context.Background())
	if err != nil {
		return nil, fmt.Errorf("cloudsqlconn.NewDialer: %w", err)
	}

	var opts []cloudsqlconn.DialOption

	mysql.RegisterDialContext("cloudsqlconn",
		func(ctx context.Context, addr string) (net.Conn, error) {
			return d.Dial(ctx, instanceConnectionName, opts...)
		})

	dsn := fmt.Sprintf("%s:%s@cloudsqlconn(localhost:3306)/%s?parseTime=true",
		dbUser, dbPwd, dbName)

	dbPool, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("sql.Open: %w", err)
	}
	return dbPool, nil
}
