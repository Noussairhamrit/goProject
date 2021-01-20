package model
type User struct {
	Id          int
	Username    string
	Password    []byte
	Vault_token string

}

//list des users
type Users []User

///constructeur
func NewUser() *User{
	return &User{}
}