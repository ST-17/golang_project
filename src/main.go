package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
	
	_ "github.com/go-sql-driver/mysql"
	
)

type Article struct{
	Id uint16
	Title, Anons, Price string
}
type User struct{
	Id uint16
	Name, Email, Password string
}

var posts = []Article{}
var showPost = Article{}
var user = User{}

func index(w http.ResponseWriter, r *http.Request){
	t, err := template.ParseFiles("templates/index.html","templates/header.html","templates/footer.html")

	if err != nil{
		fmt.Fprint(w, err.Error())
	}

	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/golang")
    if err != nil {
        panic(err)
    }
	defer db.Close()

	res, err := db.Query("SELECT * FROM `article`")
	if err != nil {
        panic(err)
    }

	posts = []Article{}
	for res.Next(){
		var post Article
		err = res.Scan(&post.Id, &post.Title, &post.Anons, &post.Price)
		if err != nil {
			panic(err)
		}

		posts = append(posts, post)
	}

	t.ExecuteTemplate(w, "index", posts)
}

func create(w http.ResponseWriter, r *http.Request){
	t, err := template.ParseFiles("templates/create.html","templates/header.html","templates/footer.html")

	if err != nil{
		fmt.Fprint(w, err.Error())
	}

	t.ExecuteTemplate(w, "create", nil)
}

func save_article(w http.ResponseWriter, r *http.Request){
	title := r.FormValue("title")
	anons := r.FormValue("anons")
	price := r.FormValue("price")

	if title == "" || anons == "" || price == ""{
		fmt.Fprintf(w, "Not all info fill in")
	} else{
		db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/golang")
    if err != nil {
        panic(err)
    }
		defer db.Close()
		
		// Установка данных
		insert, err := db.Query(fmt.Sprintf("INSERT INTO `article` (`title`,`anons`,`price`) VALUES('%s','%s','%s')", title, anons, price))
		if err != nil{
			panic(err)
		}
		defer insert.Close()

		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

func register(w http.ResponseWriter, r *http.Request){
	t, err := template.ParseFiles("templates/register.html","templates/header.html","templates/footer.html")

	if err != nil{
		fmt.Fprint(w, err.Error())
	}

	t.ExecuteTemplate(w, "register", nil)
}

func save_register(w http.ResponseWriter, r *http.Request){
	user_name := r.FormValue("user_name")
	user_email := r.FormValue("user_email")
	user_password := r.FormValue("user_password")
	user_repeat_password := r.FormValue("user_repeat_password")


	if user_name != "" && user_email != "" && user_password != "" && user_repeat_password != ""{
		db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/golang")
		if err != nil{
			panic(err)
		}
		// Установка данных
		insert, err := db.Query(fmt.Sprintf("INSERT INTO `login` (`name`,`email`,`password`) VALUES('%s','%s','%s')", user_name, user_email, user_password))
		if err != nil{
			panic(err)
		}

		defer db.Close()
		defer insert.Close()
	}
	w.WriteHeader(http.StatusNoContent)
}

func login(w http.ResponseWriter, r *http.Request){
	t, err := template.ParseFiles("templates/login.html","templates/header.html","templates/footer.html")

	if err != nil{
		fmt.Fprint(w, err.Error())
	}
	
	t.ExecuteTemplate(w, "login", nil)
}

func check_login(w http.ResponseWriter, r *http.Request) {
    FormEmail := r.FormValue("user_email")
    FormPassword := r.FormValue("user_password")

    db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/golang")
    if err != nil {
        panic(err)
    }

    res, err := db.Query(fmt.Sprintf("SELECT * FROM `login` WHERE `email` = '%s' and `password` = '%s'", FormEmail, FormPassword))
    if err != nil {
        panic(err)
    }
    
    if !res.Next() {
        w.WriteHeader(http.StatusNoContent)
    }
    
    http.Redirect(w, r, "/", http.StatusSeeOther)
	defer db.Close()
}

func show_post(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)

	t, err := template.ParseFiles("templates/show.html","templates/header.html","templates/footer.html")


    db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/golang")
    if err != nil {
        panic(err)
    }
	defer db.Close()

	res, err := db.Query(fmt.Sprintf("SELECT * FROM `article` WHERE `id` = '%s'", vars["id"]))
	if err != nil {
        panic(err)
    }

	showPost = Article{}
	for res.Next(){
		var post Article
		err = res.Scan(&post.Id, &post.Title, &post.Anons, &post.Price)
		if err != nil {
			panic(err)
		}

		showPost = post
	}

	t.ExecuteTemplate(w, "show", showPost)
}

func searchHandler(w http.ResponseWriter, r *http.Request) {
    query := r.FormValue("query")
	if query == "" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}else{
	t, err := template.ParseFiles("templates/index2.html","templates/header.html","templates/footer.html")

	if err != nil{
		fmt.Fprint(w, err.Error())
	}

	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/golang")
    if err != nil {
        panic(err)
    }
	defer db.Close()

	res, err := db.Query(fmt.Sprintf("SELECT * FROM `article` WHERE `title` = '%s'", query))
	if err != nil {
        panic(err)
    }

	posts = []Article{}
	for res.Next(){
		var post Article
		err = res.Scan(&post.Id, &post.Title, &post.Anons, &post.Price)
		if err != nil {
			panic(err)
		}

		posts = append(posts, post)
	}

	t.ExecuteTemplate(w, "index2", posts)
}
}


func handleFunc(){
	rtr := mux.NewRouter()
	rtr.HandleFunc("/", index).Methods("GET")
	rtr.HandleFunc("/create/", create).Methods("GET")
	rtr.HandleFunc("/save_article/", save_article).Methods("POST")
	rtr.HandleFunc("/register/", register).Methods("GET")
	rtr.HandleFunc("/save_register/", save_register).Methods("POST")
	rtr.HandleFunc("/login/", login).Methods("GET")
	rtr.HandleFunc("/check_login/", check_login).Methods("POST")
	rtr.HandleFunc("/post/{id:[0-9]+}", show_post).Methods("GET", "POST")
	rtr.HandleFunc("/search/", searchHandler).Methods("GET", "POST")

	http.Handle("/", rtr)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
	http.ListenAndServe(":8080", nil)
}



func main(){
	handleFunc()
}