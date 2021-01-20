package controller

import (
	"fmt"
	"net/http"
	_ "net/http"
	"github.com/gorilla/mux"
	"html/template"
	"github.com/crud_go_postgres/service"
	"github.com/crud_go_postgres/model"
	"strconv"
)
var templates *template.Template

func Global()  {
	templates=template.Must(template.ParseGlob("templates/*.html"))
	//////////////
	r:=mux.NewRouter()
	r.HandleFunc("/",indexGethandler).Methods("GET")
	r.HandleFunc("/",indexPosthandler).Methods("POST")
	////////////////////////////
	http.Handle("/",r)
	fmt.Println("server startting ...")
	http.ListenAndServe(":3000",nil)
	////////////////////////////
}
func indexGethandler(w http.ResponseWriter ,r*http.Request){
		users:=service.FindAllUser()
	templates.ExecuteTemplate(w,"index.html",users)
}
func indexPosthandler(w http.ResponseWriter ,r*http.Request){
	r.ParseForm()
	id:=r.PostForm.Get("id")
	id1,_:=strconv.ParseInt(id,10,64)
	id2:=int(id1)
	username := r.PostForm.Get("username")
	password := r.PostForm.Get("password")
	vault_token := r.PostForm.Get("vault_token")
	user:=model.User{id2,username,password,vault_token}
	service.CreateUser(user)
	http.Redirect(w,r,"/",302)
}