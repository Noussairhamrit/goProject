package test

import (
	"github.com/crud_go_postgres/model"
	"github.com/crud_go_postgres/service"
	"testing"
)
var usertest=model.User{
	Id:8,
	Username:"fr",
	Vault_token: "2200"}


func TestFindUserByID(t *testing.T)  {
	t.Run("test is ok", func(t *testing.T) {
		userExpected:=service.FindUserByID(usertest.Id)
		if usertest.Id !=  userExpected.Id ||
					usertest.Username !=  userExpected.Username ||
					usertest.Vault_token !=  userExpected.Vault_token {
			t.Fail()
		}

	})
}


