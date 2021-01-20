package model
type User struct {
	Id          int
	Username    string
	Password    string
	Vault_token string

}

//list des users
type Users []User

///constructeur
func NewUser() *User{
	return &User{}
}