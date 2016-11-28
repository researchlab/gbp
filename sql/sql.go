package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type DbWorker struct {
	Dsn        string
	Db         *sql.DB
	UserInfo   userTB
	InsertStmt *sql.Stmt
	QueryStmt  *sql.Stmt
}
type userTB struct {
	Id   int
	Name sql.NullString
	Age  sql.NullInt64
}

func main() {
	var err error
	dbw := DbWorker{
		Dsn: "root:123456@tcp(localhost:3306)/sqlx_db?charset=utf8mb4",
	}
	dbw.Db, err = sql.Open("mysql", dbw.Dsn)
	if err != nil {
		panic(err)
		return
	}
	defer dbw.Db.Close()
	if err, ok := dbw.PreWork(); ok {
		dbw.insertData()
		dbw.QueryData()
	} else {
		panic(err)
		return
	}
}

func (dbw *DbWorker) PreWork() (error, bool) {
	var err error
	if dbw.InsertStmt, err = dbw.Db.Prepare(`INSERT INTO user (name, age) VALUES (?, ?)`); nil != err {
		return err, false
	}

	if dbw.QueryStmt, err = dbw.Db.Prepare(`SELECT * From user where age >= ? AND age < ?`); nil != err {
		return err, false
	}
	return nil, true
}

func (dbw *DbWorker) insertData() {
	ret, err := dbw.InsertStmt.Exec("xys", 23)
	if err != nil {
		fmt.Printf("insert data error: %v\n", err)
		return
	}
	if LastInsertId, err := ret.LastInsertId(); nil == err {
		fmt.Println("LastInsertId:", LastInsertId)
	}
	if RowsAffected, err := ret.RowsAffected(); nil == err {
		fmt.Println("RowsAffected:", RowsAffected)
	}
}

func (dbw *DbWorker) QueryDataPre() {
	dbw.UserInfo = userTB{}
}
func (dbw *DbWorker) QueryData() {
	dbw.QueryDataPre()
	rows, err := dbw.QueryStmt.Query(20, 30)
	defer rows.Close()
	if err != nil {
		fmt.Printf("insert data error: %v\n", err)
		return
	}
	for rows.Next() {
		rows.Scan(&dbw.UserInfo.Id, &dbw.UserInfo.Name, &dbw.UserInfo.Age)
		if err != nil {
			fmt.Printf(err.Error())
			continue
		}
		if !dbw.UserInfo.Name.Valid {
			dbw.UserInfo.Name.String = ""
		}
		if !dbw.UserInfo.Age.Valid {
			dbw.UserInfo.Age.Int64 = 0
		}
		fmt.Println("get data, id: ", dbw.UserInfo.Id, " name: ", dbw.UserInfo.Name.String, " age: ", int(dbw.UserInfo.Age.Int64))
	}

	err = rows.Err()
	if err != nil {
		fmt.Printf(err.Error())
	}
}
