package main

import (
	"fmt"
	"example.com/hello/mapdb"
	// "time"
)

func main () {
	db := mapdb.NewDB()
	fmt.Printf("Yay! A new db! %v\n", db)
	mapdb.DBAdd(db, "hello", "world")
	fmt.Printf("Adding hello:world... %v\n", db)
	val := mapdb.DBGet(db, "hello")
	fmt.Printf("Value of key hello is: %v\n", val)
	fmt.Printf("This shouldn't work: %v\n", mapdb.DBGet(db, "world"))
	mapdb.DBDelete(db, "hello")
	fmt.Printf("And now DB is empty: %v\n", db)
}