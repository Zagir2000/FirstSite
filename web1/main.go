package main

import (
	"fmt"
	"html/template" //Вывод полноценных html страниц
	"net/http"
)

type User struct {
	Name                  string
	Age                   uint16
	Money                 int16
	Avg_frades, happiness float64
	Hobbies               []string
}

func (u User) getAllinfo() string {
	return fmt.Sprintf("User name is %s. She is %d and she has %d dollars", u.Name, u.Age, u.Money)
}

func (u *User) setNewName(newName string) {
	u.Name = newName
}

func home_page(w http.ResponseWriter, r *http.Request) { //http.ResponseWriter- что лио выводить на страницу, http.Request - запрос
	Nailya := User{"Nailya", 22, 560, 4.8, 0.5, []string{"Reading", "Dance"}}
	//Nailya.setNewName("Zagir")
	//fmt.Fprintf(w, `<h1>Main Text <h1>
	//<b>main</b>`)
	tmpl, _ := template.ParseFiles("templates/home_page.html")
	tmpl.Execute(w, Nailya)
}
func contacts_page(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Contacts_page")
}

func HandleRequest() {
	//Создание севера
	http.HandleFunc("/", home_page) // отследить переходи по URL адресу и при переходу на этот url адрес вызвать какой либо метод, кторый	 мы будем показывать пользователю, "/" отслеживать главную страницу, метод home_page
	http.HandleFunc("/contacts/", contacts_page)
	http.ListenAndServe(":8080", nil) //nil - ничего не передаем

}
func main() {
	HandleRequest()
}
