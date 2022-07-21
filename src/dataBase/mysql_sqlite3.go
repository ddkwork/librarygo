package dataBase

import (
	"database/sql"
	"github.com/ddkwork/librarygo/src/mycheck"

	_ "github.com/mattn/go-sqlite3"
)

type (
	Interface interface {
		Init(driverName, dataSourceName string) bool
		CreatTables(DDL string) bool
		Query(query string) (ok bool)
		QueryResult() any
		Update(query string, args ...any) bool
		Insert(query string, args ...any) bool
	}
	object struct {
		//Client *redis.Client //写到另外的文件，移除工程的爬虫工程的全局变量
		db          *sql.DB
		stmt        *sql.Stmt
		rows        *sql.Rows
		result      sql.Result
		queryResult any
		err         error
	}
)

func New() Interface { return &object{} }

var (
	Default = New()
)

func (o *object) Init(driverName, dataSourceName string) bool {
	o.db, o.err = sql.Open(driverName, dataSourceName)
	if !mycheck.Error(o.err) {
		return false
	}
	o.db.SetMaxOpenConns(1000)
	o.db.SetMaxIdleConns(30000)
	return mycheck.Error(o.db.Ping())
}
func (o *object) CreatTables(DDL string) bool { return mycheck.Error2(o.db.Exec(DDL)) }
func (o *object) QueryResult() any            { return o.queryResult }
func (o *object) Query(query string) (ok bool) {
	o.rows, o.err = o.db.Query(query)
	if !mycheck.Error(o.err) {
		return
	}
	defer func() {
		if o.rows == nil {
			mycheck.Error("rows == nil ")
			return
		}
		mycheck.Error(o.rows.Close())
	}()
	for o.rows.Next() {
		if !(mycheck.Error(o.rows.Scan(&o.queryResult))) {
			return
		}
	}
	return true
}

func (o *object) Update(query string, args ...any) bool { return o.stmtExec(query, args) }
func (o *object) Insert(query string, args ...any) bool { return o.stmtExec(query, args) }
func (o *object) stmtExec(query string, args ...any) bool {
	o.stmt, o.err = o.db.Prepare(query)
	if !mycheck.Error(o.err) {
		return false
	}
	defer func() {
		if o.stmt == nil {
			mycheck.Error("stmt == nil ")
			return
		}
		mycheck.Error(o.stmt.Close())
	}()
	o.result, o.err = o.stmt.Exec(args)
	return mycheck.Error(o.err)
}
