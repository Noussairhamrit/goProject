package main

import (
	"github.com/crud_go_postgres/service"
	"github.com/crud_go_postgres/controller"
)

func main(){
	service.Create_table()
	//user:=model.User{
	//	Id:          2,
	//	Username:    "azerty",
	//	Password:    "azerty",
	//	Vault_token: "azerty",
	//}
	//service.CreateUser(user)
	service.FindUserByID(2)
	//service.DeleteUser(3)
	service.EditUser(2,"nou","nou","774411")
	service.FindAllUser()
	controller.Global()
}