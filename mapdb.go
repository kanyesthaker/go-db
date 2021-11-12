package mapdb

// Extremely small database, pretty much the most naïve thing ever 

import (
	"fmt"
)

type DB map[string]interface{}

func NewDB () DB {
	db := DB{}
	return db
}

func DBAdd (db DB, key string, value interface{}) {
	db[key] = value
}

func DBGet (db DB, key string) interface{} {
	if val, ok := db[key]; ok {
		return val
	} 
	return nil
}

func DBDelete (db DB, key string) {
	if _, ok := db[key]; ok {
		delete(db, key)
	}
}
