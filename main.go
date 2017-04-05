package main

import (
	"LemonChat/database"
	"time"
	"strconv"
	"LemonChat/clients"
)

func main() {
	//msgServer := msgServer.NewMessageServer("", "" ,"localhost:1024")
	//msgServer.Serve()

	cdb, _ := database.NewDB()
	for i := 0; i < 10; i++ {
		name := "lemon" + strconv.Itoa(i)
		addr := "addr-" + strconv.Itoa(i)
		p := &clients.Person{NickName: name, Age: 30 + i, Sex: 1, Address: "cd", Avatar: addr}
		cdb.InsertPerson(p)
	}
	cdb.Query()
	time.Sleep(1000)
}
