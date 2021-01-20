package service

import (
	//"database/sql"
	"fmt"
	"github.com/crud_go_postgres/config"
	"github.com/crud_go_postgres/model"
	"log"

	//"github.com/crud_go_postgres/repository"
	//"log"
)
var db=config.OpenConnection()
func Create_table()  {

	_,err:=db.Exec("CREATE TABLE IF NOT EXISTS users ( id integer ,username varchar(255) , password varchar(255) , vault_token varchar(255) )")

	if err != nil {
		panic(err)
	} else {
		fmt.Println("table user created")
	}

}
func FindAllUser() []model.User {
	rows, err := db.Query("SELECT * FROM \"users\"")
	if err != nil {
		return nil
	}else {
		var users []model.User
		for rows.Next(){
			var id int
			var username string
			var password string
			var vault_token string
			err2 :=rows.Scan(&id,&username,&password,&vault_token)
			if err2 != nil {
				return users
			}else{
				user:=model.User{
					Id:          id,
					Username:    username,
					Password:    password,
					Vault_token: vault_token,
				}
				users=append(users ,user)

			}

		}
		fmt.Println("list of all user",users)
		return users

	}

}
func CreateUser(user model.User)  {
	insertDynStmt := `insert into "users" ("id", "username","password","vault_token") values($1, $2 , $3, $4)`

	_, e := db.Exec(insertDynStmt, user.Id, user.Username, user.Password, user.Vault_token)
	CheckError(e)

}

func FindUserByID (id int) model.User {
	var user model.User
	rows, err := db.Query(`SELECT * FROM users where id=$1`, id)
	if err != nil {
		log.Fatal(err)
	} else {

		if rows.Next() {
			var id int
			var username string
			var password string
			var vault_token string
			err2 := rows.Scan(&id, &username, &password, &vault_token)
			if err2 != nil {
				return user
			} else {
				user = model.User{
					Id:          id,
					Username:    username,
					Password:    password,
					Vault_token: vault_token,
				}


			}

		}

	}

	fmt.Println("find user by ID",user)
	return user
}
func DeleteUser(id int)  {
		DeleteStmt := `DELETE FROM "users" WHERE id=$1`

	_, e := db.Exec(DeleteStmt, id)
	CheckError(e)
	fmt.Println("user with id :",id," is deleted")

}
func EditUser(id int,newusername string, newpassword string,newvault_token string)  {
	EditStmt := `UPDATE "users" SET username=$1 ,password=$2, vault_token=$3 WHERE id=$4`
	_, e := db.Exec(EditStmt,newusername, newpassword,newvault_token,id)
	CheckError(e)
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}