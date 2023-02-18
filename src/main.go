package main

import(
	"fmt"
	"net/http"
	"html/template"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func index(w http.ResponseWriter, r *http.Request){
	t, err := template.ParseFiles("templates/index.html","templates/header.html","templates/footer.html")

	if err != nil{
		fmt.Fprint(w, err.Error())
	}

	t.ExecuteTemplate(w, "index", nil)
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
	} else {
		db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/goland")
		if err != nil{
			panic(err)
		}
		
		// Установка данных
		insert, err := db.Query(fmt.Sprintf("INSERT INTO `article` (`title`,`anons`,`price`) VALUES('%s','%s','%s')", title, anons, price))
		if err != nil{
			panic(err)
		}

		http.Redirect(w, r, "/", http.StatusSeeOther)
		defer db.Close()
		defer insert.Close()
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


	if user_name == "" || user_email == "" || user_password == "" || user_repeat_password == ""{
		fmt.Fprintf(w, "Not all info fill in")
	} else {
		db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/goland")
		if err != nil{
			panic(err)
		}
		
		// Установка данных
		insert, err := db.Query(fmt.Sprintf("INSERT INTO `login` (`name`,`email`,`password`) VALUES('%s','%s','%s')", user_name, user_email, user_password))
		if err != nil{
			panic(err)
		}

		http.Redirect(w, r, "/", http.StatusSeeOther)
		defer db.Close()
		defer insert.Close()
	}
}

func handleFunc(){
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
	http.HandleFunc("/", index)
	http.HandleFunc("/create/", create)
	http.HandleFunc("/save_article/", save_article)
	http.HandleFunc("/register/", register)
	http.HandleFunc("/save_register/", save_register)
	http.ListenAndServe(":8080", nil)
}

func main(){
	handleFunc()
}