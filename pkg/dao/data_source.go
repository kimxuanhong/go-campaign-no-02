package dao

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type DataSource struct {
	db *sql.DB
}

var instanceDataSource *DataSource

func NewDataSource() *DataSource {
	if instanceDataSource == nil {
		instanceDataSource = &DataSource{}
		instanceDataSource.Init()
	}
	return instanceDataSource
}

func (r *DataSource) GetConn() *sql.DB {
	return r.db
}

func (r *DataSource) Init() {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "passw0rd"
	dbName := "golang_demo"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		fmt.Println("Connection Failed!!")
	}

	err = db.Ping()
	if err != nil {
		fmt.Println("Ping Failed!!")
	}

	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(time.Second * 10)

	r.db = db
}

//CloseStmt after run stmt
func (r *DataSource) CloseStmt(stmt *sql.Stmt) {
	if stmt != nil {
		err := stmt.Close()
		if err != nil {
			return
		}
	}
}

//CloseRows when select
func (r *DataSource) CloseRows(rows *sql.Rows) {
	if rows != nil {
		err := rows.Close()
		if err != nil {
			return
		}
	}
}
