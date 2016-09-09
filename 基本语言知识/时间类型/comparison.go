package main

import (
	"fmt"
	"reflect"
	"time"
)

func main() {
	// utc life
	loc, _ := time.LoadLocation("UTC")

	// setup a start and end time
	createdAt := time.Now().In(loc).Add(1 * time.Hour)
	expiresAt := time.Now().In(loc).Add(4 * time.Hour)

	// get the diff
	diff := expiresAt.Sub(createdAt)
	t := reflect.TypeOf(diff)

	fmt.Println("type", t.Name())
	fmt.Println("Lifespan is %+v", diff)
}
