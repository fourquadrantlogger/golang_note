package main

import (
	"gopkg.in/mgo.v2"
	"log"
)

func insertmap(session *mgo.Session, clt *mgo.Collection) {
	_row := map[string]interface{}{
		"myd":      234,
		"username": map[string]interface{}{"baba": 123},
	}
	if err := clt.Insert(_row); err != nil {
		log.Printf("insert %v failed: %v\n", _row, err)
	}

}
func main() {
	session, err := mgo.Dial("127.0.0.1:27017")
	if err != nil {
		log.Fatal("无法打开MongoDB！")
		return
	}
	defer session.Close()
	clt := session.DB("test").C("test")
	//Collontion(session,clt)
	insertmap(session, clt)
}
