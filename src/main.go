package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	_ "github.com/go-sql-driver/mysql"
)

type Article struct{
	Id uint16
	Title, Anons string
	Price int
}
type Review struct{
	Id uint16
	Name, Text string
	Rating int
}
type User struct{
	Id uint16
	Name, Email, Password string
}

var posts = []Article{}
var comments = []Review{}
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
	price, err := strconv.Atoi(r.FormValue("price"))
	
	if err != nil {
		fmt.Println("Error during conversion")
	}

	if title == "" || anons == ""{
		fmt.Fprintf(w, "Not all info fill in")
	} else{
		db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/golang")
    if err != nil {
        panic(err)
    }
		defer db.Close()
		
		insert, err := db.Query(fmt.Sprintf("INSERT INTO `article` (`title`,`anons`,`price`) VALUES('%s','%s','%d')", title, anons, price))
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

func filterHandler(w http.ResponseWriter, r *http.Request) {
    query1 := r.FormValue("query1")
	query2 := r.FormValue("query2")
	
	if query1 == "" || query2== ""{
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

	res, err := db.Query(fmt.Sprintf("SELECT * FROM `article` WHERE `price` >= '%s' AND `price` <= '%s'", query1, query2))
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

func review(w http.ResponseWriter, r *http.Request){
	t, err := template.ParseFiles("templates/review.html","templates/header.html","templates/footer.html")

	if err != nil{
		fmt.Fprint(w, err.Error())
	}

	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/golang")
    if err != nil {
        panic(err)
    }
	defer db.Close()

	res, err := db.Query("SELECT * FROM `review`")
	if err != nil {
        panic(err)
    }

	comments = []Review{}
	for res.Next(){
		var comment Review
		err = res.Scan(&comment.Id, &comment.Name, &comment.Text, &comment.Rating)
		if err != nil {
			panic(err)
		}

		comments = append(comments, comment)
	}

	t.ExecuteTemplate(w, "review", comments)
}

func save_review(w http.ResponseWriter, r *http.Request){
	name := r.FormValue("name")
	text := r.FormValue("text")
	rating, err := strconv.Atoi(r.FormValue("rating"))
	
	if err != nil {
		fmt.Println("Error during conversion")
	}

	if name == "" || text == ""{
		fmt.Fprintf(w, "Not all info fill in")
	} else{
		db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/golang")
    if err != nil {
        panic(err)
    }
		defer db.Close()
		
		insert, err := db.Query(fmt.Sprintf("INSERT INTO `review` (`name`,`text`,`rating`) VALUES('%s','%s','%d')", name, text, rating))
		if err != nil{
			panic(err)
		}
		defer insert.Close()

		http.Redirect(w, r, "/review/", http.StatusSeeOther)
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
	rtr.HandleFunc("/filter/", filterHandler).Methods("GET", "POST")
	rtr.HandleFunc("/review/", review).Methods("GET")
	rtr.HandleFunc("/save_review/", save_review).Methods("POST")

	http.Handle("/", rtr)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
	http.ListenAndServe(":8080", nil)
}



func main(){
	handleFunc()
}