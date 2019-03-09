package services

import (
	"go-server/model"
	"go-server/repositories"
)

//
//var p = fmt.Println
//
//type JSONString string
//
//func (j JSONString) MarshalJSON() ([]byte, error) {
//	return []byte(j), nil
//}

type TestService struct {
	Name string "test"
}

var Services TestService

func (self TestService) Test()(t333 string){
	//fmt.Print("33333")
	//return TestService{.}
	//bytes.ToTitle(RandString()
	t333 = "rerere"
	return
}

func RandString() (string){
	//fmt.Print(2222)
	// Pretend to return a random string
	return "testetestestestetsetstes"
	//return string(Services.Name)
}

func MemberSelect() []model.Member{


	return repositories.Repo.Select()
	//return "memberselect"
	//return []model.Member
}
