package main

import (
	"time"
	"fmt"
)
var (sleeptime=time.Second*1)
type Msg struct{
	chanid int
	data string
}

func Sender(msg chan *Msg){
	count:=0
	for true{
		count++
		time.Sleep(sleeptime)
		msg<-&Msg{count%5,time.Now().String()[0:19]}
	}
}

func Receiver(chanid int,msg chan *Msg)  {
	for true {
		m := <-msg

		fmt.Println(chanid,*m)
	}
}
func main() {
	Message:=make(chan *Msg)
	go Sender(Message)
	go Receiver(0,Message)
	go Receiver(1,Message)
	go Receiver(2,Message)
	go Receiver(3,Message)
	go Receiver(4,Message)

	for true{}
}
