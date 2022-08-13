package dao

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

//go:generate mockgen -source=data_source.go -destination=mocks/data_source_mock.go -package=mocks

type DataSource interface {
	init() *sql.DB
	GetConn() *sql.DB
	CloseStmt(stmt *sql.Stmt)
	CloseRows(rows *sql.Rows)
}

type DataSourceImpl struct {
	db *sql.DB
}

var instanceDataSource *DataSourceImpl

func NewDataSource() *DataSourceImpl {
	if instanceDataSource == nil {
		instanceDataSource = &DataSourceImpl{}
		instanceDataSource.init()
	}
	return instanceDataSource
}

func (r *DataSourceImpl) GetConn() *sql.DB {
	if r.db == nil {
		r.db = r.init()
	}
	return r.db
}

func (r *DataSourceImpl) init() *sql.DB {
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

	return db
}

//CloseStmt after run stmt
func (r *DataSourceImpl) CloseStmt(stmt *sql.Stmt) {
	if stmt != nil {
		err := stmt.Close()
		if err != nil {
			return
		}
	}
}

//CloseRows when select
func (r *DataSourceImpl) CloseRows(rows *sql.Rows) {
	if rows != nil {
		err := rows.Close()
		if err != nil {
			return
		}
	}
}
