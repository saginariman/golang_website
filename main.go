package main
import (
	"fmt" 
	"net/http" 
	"html/template" 
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	// "os"
)

type Article struct {
	Id uint16
	Title, Anons, FullText string
}

var articles = []Article{}
var showPost = Article{}

func index (w http.ResponseWriter, r *http.Request){
	t, err := template.ParseFiles("templates/index.html", "templates/header.html", "templates/footer.html")

	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	db, err := sql.Open("mysql", "mysql:mysql@tcp(127.0.0.1:3306)/golang")
	if err != nil {
		panic(err)
	}

	defer db.Close()
	
	// Выборка данных
	res, err := db.Query("SELECT * FROM `articles`")
	if err != nil{
		panic(err)
	}

	articles = []Article{}
	for res.Next() {
		var article Article
		err = res.Scan(&article.Id, &article.Title, &article.Anons, &article.FullText)
		if err != nil {
			panic(err)
		}

		fmt.Println(fmt.Sprintf("Article: %d %s", article.Id, article.Title))
		articles = append(articles, article)
	}

	t.ExecuteTemplate(w, "index",  articles)
}
func create (w http.ResponseWriter, r *http.Request){
	t, err := template.ParseFiles("templates/create.html", "templates/header.html", "templates/footer.html")

	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	t.ExecuteTemplate(w, "create",  nil)
}

func save_article (w http.ResponseWriter, r *http.Request){
	title := r.FormValue("title")
	anons := r.FormValue("anons")
	full_text := r.FormValue("full_text")

	if title == "" || anons == "" || full_text == "" {
		fmt.Fprintf(w, "Не все данные заполнены")
		// os.Exit(1)
	}else {
		db, err := sql.Open("mysql", "mysql:mysql@tcp(127.0.0.1:3306)/golang")
		if err!=nil {
			panic(err)
		}

		defer db.Close()
		
		// Установка данных
		insert, err := db.Query(fmt.Sprintf("INSERT INTO `articles` (`title`, `anons`, `full_text`) VALUES('%s', '%s', '%s')", title, anons, full_text))
		if err != nil {
			panic(err)
		}
		defer insert.Close()

		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

func show_article(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	t, err := template.ParseFiles("templates/show.html", "templates/header.html", "templates/footer.html")

	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	// w.WriterHeader(http.StatusOK)
	// fmt.Fprintf(w, "id is %s", vars["id"])
	db, err := sql.Open("mysql", "mysql:mysql@tcp(127.0.0.1:3306)/golang")
	if(err != nil){
		panic(err)
	}
	defer db.Close()
	// Выборка данных
	res, err := db.Query(fmt.Sprintf("SELECT * FROM `articles` WHERE `id` = '%s'", vars["id"]))
	if err != nil{
		panic(err)
	}

	showPost = Article{}
	for res.Next() {
		var article Article
		err = res.Scan(&article.Id, &article.Title, &article.Anons, &article.FullText)
		if err != nil {
			panic(err)
		}

		showPost = article
	}

	t.ExecuteTemplate(w, "show",  showPost)
}

func handleFunc() {
	rtr := mux.NewRouter()
	rtr.HandleFunc("/", index).Methods("GET")
	rtr.HandleFunc("/create", create).Methods("GET")
	rtr.HandleFunc("/save_article", save_article).Methods("POST")
	rtr.HandleFunc("/post/{id:[0-9]+}",  show_article).Methods("GET")

	http.Handle("/", rtr)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
	http.ListenAndServe(":8080", nil)
}

func main() {
	handleFunc()
}