package repositories

import (
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
	"go-server/model"
	"go-server/servers"
)

type Repositories struct {

}

//type Member struct {
//	No int64 `db:"member_no"`
//	Name string `db:"member_name"`
//	Id string `db:"member_id"`
//	RegDate string `db:"member_reg_date"`
//	UpdDate string `db:"member_upd_date"`
//}

var Repo Repositories

func (self Repositories) Select() [] model.Member {
	//servers.DbServers.Connect2()
	//
	//m := servers.DbServers.InitDb()
	//
	//count, err := m.SelectInt("select count(*) from member")
	//servers.CheckErr(err, "select count(*) failed")
	//log.Println("Rows after inserting:", count)
	//
	//var members []model.Member
	//_, err = m.Select(&members, "select * from member")
	//servers.CheckErr(err, "Select failed")
	//return members

	db := servers.Database().GetInstance()
	var members []model.Member
	db.Select(&members, "select * from member")
	return members
}