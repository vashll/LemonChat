package database

import (
	"testing"
	"fmt"
	"LemonChat/clients"
)

func TestNewDB(t *testing.T) {
	_, err := NewDB()
	if err != nil {
		fmt.Println("New db failed!")
	}
}

func TestChatDB_InsertPerson(t *testing.T) {
	db,_ := NewDB()
	p := &clients.Person{NickName:"lemon", Age:30, Sex:1, Address:"cd", Avatar:""}
	//p2 := clients.Person{NickName:"Leena", Age:26, Sex:0, Address:"cd", Avatar:""}

	db.InsertPerson(p)
	//db.InsertPerson(p2)
}
