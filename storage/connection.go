package storage

import (
	"database/sql"
	"log"
	"sync"

	// !!
	_ "github.com/go-sql-driver/mysql"
)

var (
	once     sync.Once
	instance *ConnectionMysql
)

// ConnectionMysql is the class of connection
type ConnectionMysql struct {
	db *sql.DB
}

// NewConnection is the constructor and genered single instace
func NewConnection() *ConnectionMysql {
	once.Do(func() {
		db, err := sql.Open("mysql", "leo:chester@tcp(localhost:3306)/BD_REST_GO")
		if err != nil {
			log.Fatalf("HA OCURRIDO UN ERROR AL INTENTAR ACCEDER A LA BD: %+v", err)
		}
		if err = db.Ping(); err != nil {
			log.Fatalf("HA OCURRIDO UN ERROR EN EL PING: %+v", err)
		}
		instance = createSingleInstance(db)
	})
	return instance
}

func createSingleInstance(db *sql.DB) *ConnectionMysql {
	return &ConnectionMysql{db}
}

// Pool method of the class ConnectionMysql
func (mysql *ConnectionMysql) Pool() *sql.DB {
	return mysql.db
}
