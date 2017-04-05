package database

import (
	"database/sql"
	//"os"
	"log"
	"LemonChat/clients"
	"fmt"
	_"github.com/go-sql-driver/mysql"
)

const (
	DB_NAME      string = "chatdb"
	PERSON_TABLE string = "person"
)

type ChatDB struct {
	DB *sql.DB
}

func NewDB() (*ChatDB, error) {
	datasource := "root:lemon@/" + DB_NAME
	db, err := sql.Open("mysql", datasource)
	checkErr(err)
	if db == nil {
		log.Println("DB connection error")
		return nil, err
		//os.Exit(0)
	}
	db.SetMaxOpenConns(500)
	db.SetMaxIdleConns(200)
	db.Ping()
	return &ChatDB{DB: db}, nil
}

func (this *ChatDB) InsertPerson(p *clients.Person) error {
	sqlstr := "INSERT INTO " + PERSON_TABLE + " (nickname, sex, age, address, avatar_url) VALUES (?, ?, ?, ?, ?)"
	stmt, err := this.DB.Prepare(sqlstr)
	defer stmt.Close()
	checkErr(err)
	_, e := stmt.Exec(p.NickName, p.Sex, p.Age, p.Address, p.Avatar)
	return e
}

func (this *ChatDB) Query() {
	rows, _ := this.DB.Query("SELECT * FROM person")
	fmt.Println(rows.Columns())
	defer rows.Close()
	for rows.Next() {
		var id int
		var name string
		var sex int
		var age int
		var addr string
		var ava string
		if err := rows.Scan(&id, &name, &sex, &age, &addr, &ava); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%d", id)
		fmt.Printf("%s", " ")
		fmt.Printf("%s", name)
		fmt.Printf("%s", " ")
		fmt.Printf("%d", sex)
		fmt.Printf("%s", " ")
		fmt.Printf("%d", age)
		fmt.Printf("%s", " ")
		fmt.Printf("%s", addr)
		fmt.Printf("%s", " ")
		fmt.Printf("%s", ava)
		fmt.Println()
	}
}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
		return
	}
}
