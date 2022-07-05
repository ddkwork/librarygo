package dataBase

import (
	"database/sql"
	"github.com/ddkwork/librarygo/src/check"

	_ "github.com/mattn/go-sqlite3"
)

type (
	Interface interface {
		Init(driverName, dataSourceName string) bool
		CreatTables(DDL string) bool
		Query(query string) (ok bool)
		QueryResult() interface{}
		Update(query string, args ...interface{}) bool
		Insert(query string, args ...interface{}) bool
	}
	object struct {
		//Client *redis.Client //写到另外的文件，移除工程的爬虫工程的全局变量
		db          *sql.DB
		stmt        *sql.Stmt
		rows        *sql.Rows
		result      sql.Result
		queryResult interface{}
		err         error
	}
)

func New() Interface { return &object{} }

var (
	Default = New()
)

func (o *object) Init(driverName, dataSourceName string) bool {
	o.db, o.err = sql.Open(driverName, dataSourceName)
	if !check.Error(o.err) {
		return false
	}
	o.db.SetMaxOpenConns(1000)
	o.db.SetMaxIdleConns(30000)
	return check.Error(o.db.Ping())
}
func (o *object) CreatTables(DDL string) bool { return check.Error2(o.db.Exec(DDL)) }
func (o *object) QueryResult() interface{}    { return o.queryResult }
func (o *object) Query(query string) (ok bool) {
	o.rows, o.err = o.db.Query(query)
	if !check.Error(o.err) {
		return
	}
	defer func() {
		if o.rows == nil {
			check.Error("rows == nil ")
			return
		}
		check.Error(o.rows.Close())
	}()
	for o.rows.Next() {
		if !(check.Error(o.rows.Scan(&o.queryResult))) {
			return
		}
	}
	return true
}

func (o *object) Update(query string, args ...interface{}) bool { return o.stmtExec(query, args) }
func (o *object) Insert(query string, args ...interface{}) bool { return o.stmtExec(query, args) }
func (o *object) stmtExec(query string, args ...interface{}) bool {
	o.stmt, o.err = o.db.Prepare(query)
	if !check.Error(o.err) {
		return false
	}
	defer func() {
		if o.stmt == nil {
			check.Error("stmt == nil ")
			return
		}
		check.Error(o.stmt.Close())
	}()
	o.result, o.err = o.stmt.Exec(args)
	return check.Error(o.err)
}
