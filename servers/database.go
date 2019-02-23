package servers

import (
	"database/sql"
	"fmt"
	"log"
	_ "github.com/go-sql-driver/mysql"
)

type DbServer struct {

}

var DbServers DbServer

func (self DbServer) Connect() {
	// sql.DB 객체 생성
	db, err := sql.Open("mysql", "root:password@/demo")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// 하나의 Row를 갖는 SQL 쿼리
	var name string
	err = db.QueryRow("SELECT id FROM todo_models WHERE id = 1").Scan(&name)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(name)
}
