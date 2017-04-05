package protocols

import (
	"testing"
	"log"
)

func TestDepack(t *testing.T) {
	str := "abcd"
	b := []byte(str)
	ret := Enpack(b)
	log.Println(ret, "length = ", len(ret))
	log.Println(string(ret))
}

func TestIntToBytes(t *testing.T) {
	b := IntToBytes(8)
	if len(b) == 4 {
		log.Println("test pass",string(b))
	} else {
		t.Error("int to byte err")
	}
}
