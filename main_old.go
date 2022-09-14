package main
// import ("fmt"; "net/http"; "html/template"; "database/sql"; _ "github.com/go-sql-driver/mysql")
import ("fmt"; "database/sql"; _ "github.com/go-sql-driver/mysql")

// type User struct {
// 	Name string
// 	Age uint16
// 	Money int16
// 	Avg_grades, Happiness float64
// 	Hobbies []string
// }

// func (u User) getAllInfo() string {
// 	return fmt.Sprintf("User name is %s. He is %d and he has money "+
// 						" equal: %d", u.Name, u.Age, u.Money)
// }

// func (u *User) setNewName(newName string) {
// 	u.Name = newName
// }

// func home_page(w http.ResponseWriter, r *http.Request) {
// 	bob := User{"Bob", 25, -50, 4.2, 0.8, []string{"Football", "Skate", "Dance"}}
// 	// bob.name = "Alex"
// 	bob.setNewName("MegaMozg")
// 	// fmt.Fprintf(w, "User name is " + bob.name)
// 	// fmt.Fprintf(w, bob.getAllInfo())
// 	tmpl, _ := template.ParseFiles("templates/home_page.html")
// 	tmpl.Execute(w, bob)
// }
// func contacts_page(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Contact page!")
// }
// func handleRequest() {
// 	http.HandleFunc("/", home_page)
// 	http.HandleFunc("/contacts/", contacts_page)
// 	http.ListenAndServe(":8080", nil)
// 	// if  err == http.ErrServerClosed {
// 	// 	fmt.Println("Server started")
// 	// }
// }

type User struct {
	Id int
	Name string `json:"name"`
	Age uint16 `json:"age"`
}

func main () {
	// var bob User = ...
	// bob := User{name: "Bob", age: 25, money: -50, avg_grades: 4.2, happiness: 0.8}
	// bob := User{"Bob", 25, -50, 4.2, 0.8}
	// handleRequest()
	db, err := sql.Open("mysql", "mysql:mysql@tcp(127.0.0.1:3306)/golang")
	if err != nil {
		panic(err)
	}

	defer db.Close()

	// Установка данных
	// insert, err := db.Query("INSERT INTO `users`(`name`, `age`) VALUES('Bob', 35)")
	// if err!=nil {
	// 	panic(err)
	// }
	// defer insert.Close()
	
	// Выборка данных
	res, err := db.Query("SELECT * FROM `users`")
	if err != nil{
		panic(err)
	}

	for res.Next() {
		var user User
		err = res.Scan(&user.Id, &user.Name, &user.Age)
		if err != nil {
			panic(err)
		}

		fmt.Println(fmt.Sprintf("User: %s with age %d", user.Name, user.Age))
	}

}