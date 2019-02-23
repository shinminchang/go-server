package repositories

import "go-server/servers"

type Repositories struct {

}

var Repo Repositories

func (self Repositories) Select(){
	servers.DbServers.Connect()
}