package dao

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"os"
	"time"
)

//go:generate mockgen -source=data_source.go -destination=mocks/data_source_mock.go -package=mocks

type DataSource interface {
	GetConn() *sql.DB
	CloseStmt(stmt *sql.Stmt)
	CloseRows(rows *sql.Rows)
}

type DataSourceImpl struct {
	db *sql.DB
}

var instanceDataSource *DataSourceImpl

func DataSourceInstance() *DataSourceImpl {
	if instanceDataSource == nil {
		instanceDataSource = &DataSourceImpl{
			db: connectDB(),
		}
	}
	return instanceDataSource
}

func (r *DataSourceImpl) GetConn() *sql.DB {
	if r.db == nil {
		r.db = connectDB()
	}
	return r.db
}

func connectDB() *sql.DB {
	dbDriver := os.Getenv("DB_DRIVER")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")

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
