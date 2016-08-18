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

func Sender(msg chan Msg){
	count:=0
	for true{
		count++
		time.Sleep(sleeptime)
		msg<-Msg{count,time.Now().String()[0:19]}
	}
}

func Receiver(msg chan Msg)  {
	for true {
		m := <-msg
		fmt.Println(m)
	}
}
func main() {
	Message:=make(chan Msg)
	go Sender(Message)
	go Receiver(Message)

	for true{}
}
