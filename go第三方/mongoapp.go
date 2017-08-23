package main

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	"math/rand"
	"sync"
)

func init() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
}

func getNextSeq(session *mgo.Session) int {
	counter := session.DB("dbname").C("counter")
	cid := "counterid"

	change := mgo.Change{
		Update:    bson.M{"$inc": bson.M{"seq": 1}},
		Upsert:    true,
		ReturnNew: true,
	}
	doc := struct{ Seq int }{}
	if _, err := counter.Find(bson.M{"myd": cid}).Apply(change, &doc); err != nil {
		panic(fmt.Errorf("get counter failed:", err.Error()))
	}
	log.Println("seq:", doc)
	return doc.Seq

}
func main() {
	session, err := mgo.Dial("127.0.0.1:27017")
	if err != nil {
		log.Fatal("无法打开MongoDB！")
		return
	}
	defer session.Close()

	clt := session.DB("test").C("test")

	wg := sync.WaitGroup{}
	// 创建10个 go routine模拟多线程环境
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			_row := map[string]interface{}{
				"myd":       getNextSeq(session),
				"username":  fmt.Sprintf("name%2d", i),
				"telephone": fmt.Sprintf("tel%4d", rand.Int63n(1000))}
			if err := clt.Insert(_row); err != nil {
				log.Printf("insert %v failed: %v\n", _row, err)
			} else {
				log.Printf("insert: %v success!\n", _row)
			}
			wg.Done()
		}(i)
	}
	wg.Wait()
}
