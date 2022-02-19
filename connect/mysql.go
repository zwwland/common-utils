package connect

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type MysqlConfig struct {
	Host     string
	Port     string
	DB       string
	User     string
	Password string
	Charset  string
}

func NewMysql(rc *MysqlConfig) (*sql.DB, error) {
	mysqlUri := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s",
		rc.User,
		rc.Password,
		rc.Host,
		rc.Port,
		rc.DB,
		rc.Charset,
	)
	var conn *sql.DB
	var err error
	if conn, err = sql.Open("mysql", mysqlUri); err != nil {
		panic(err)
	}
	// conn.SetMaxIdleConns(10)
	// conn.SetMaxOpenConns(9)
	return conn, nil
}
