package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"

	_ "github.com/proullon/ramsql/driver"
)

func initDB() *sql.DB {
	db, err := sql.Open("ramsql", "TestLoadUserAddresses")
	if err != nil {
		log.Fatal(err)
	}
	res, err := db.Exec("CREATE TABLE user (id BIGSERIAL PRIMARY KEY, username TEXT, password TEXT);")
	fmt.Println("1", res, err)
	res, err = db.Exec("INSERT INTO user (username, password) VALUES ('admin', 'root');")
	fmt.Println("2", res, err)
	res, err = db.Exec("INSERT INTO user (username, password) VALUES ('sib-coder', '123123');")
	fmt.Println("3", res, err)
	res, err = db.Exec("INSERT INTO user (username, password) VALUES ('lider', 'roman');")
	fmt.Println("4", res, err)
	return db
}
func SelectFromDB() {
	db := initDB()
	pas := "root"
	user := "admin"
	row := db.QueryRow("SELECT id FROM user WHERE username=" + user + " AND password=" + pas + ";")
	var id int
	if err := row.Scan(&id); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("ID:", id)
}

func main() {
	SelectFromDB()
	db := initDB()
	fmt.Println(db)

	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		username := r.FormValue("username")
		password := r.FormValue("password")

		fmt.Println(username, " ", password)
		querySlave := "SELECT id FROM user WHERE username=" + username + " AND password=" + password + ";"
		row := db.QueryRow(querySlave)
		var id int
		if err := row.Scan(&id); err != nil {
			fmt.Println(err)
			w.Write([]byte("wrong credentials"))
			return
		}

		w.Write([]byte("ok. your id is " + strconv.Itoa(id)))
	})

	http.ListenAndServe(":8099", nil)
}

// http://localhost:8099/login?username=sib-coder&password=123123
// http://localhost:8099/login?username=admin&password="OR"1"="1