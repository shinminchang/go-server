package servers

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
	"gopkg.in/gorp.v1"
	"log"
	"sync"
)

type DbServer struct {
	//No int64 `db:"member_no"`
	//Name string `db:"member_name"`
	//Id string `db:"member_id"`
	//RegDate string `db:"member_reg_date"`
	//UpdDate string `db:"member_upd_date"`
}

var DbServers DbServer

// mysql
func (self DbServer) Connect() {
		db, err := sql.Open("mysql", "root:password@/demo")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

		var name string
	err = db.QueryRow("SELECT id FROM todo_models WHERE id = 1").Scan(&name)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(name)
}

// gorp
func (self DbServer) Connect2() {
	dbmap := DbServers.InitDb()
	defer dbmap.Db.Close()
}

func (self DbServer) InitDb() *gorp.DbMap {
			db, err := sql.Open("mysql", "root:password@/testdb")
	CheckErr(err, "sql.Open failed")

		dbmap := &gorp.DbMap{Db: db, Dialect: gorp.SqliteDialect{}}

			
			err = dbmap.CreateTablesIfNotExists()
	CheckErr(err, "Create tables failed")

	return dbmap
}

func CheckErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}

type database struct {
	db *gorp.DbMap
}

var instance *database
var once sync.Once

func Database() *database{
	once.Do(func() {
		db, err := sql.Open("mysql", "root:password@/testdb")
		CheckErr(err, "sql.Open failed")
		dbMap := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{"InnoDb", "UTF8"}}
		instance = &database{db:dbMap}
	})
	return instance
}

func (db *database) GetInstance() *gorp.DbMap{
	return db.db
}