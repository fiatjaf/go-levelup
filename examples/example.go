package levelup

import (
	"fmt"

	"github.com/fiatjaf/levelup"
)

func Example(db levelup.DB) {
	fmt.Println("setting key1 to x")
	db.Put("key1", "x")
	res, _ := db.Get("key1")
	fmt.Println("setting key2 to 2")
	fmt.Println("res at key2: ", res)
	db.Put("key2", "y")
	res, _ = db.Get("key2")
	fmt.Println("res at key2: ", res)
	fmt.Println("deleting key1")
	db.Del("key1")
	res, _ = db.Get("key1")
	fmt.Println("res at key1: ", res)

	fmt.Println("batch")
	db.Batch([]levelup.Operation{
		levelup.OpPut("key2", "w"),
		levelup.OpPut("key3", "z"),
		levelup.OpDel("key1"),
		levelup.OpPut("key1", "t"),
		levelup.OpPut("key4", "m"),
		levelup.OpPut("key5", "n"),
		levelup.OpDel("key3"),
	})
	res, _ = db.Get("key1")
	fmt.Println("res at key1: ", res)
	res, _ = db.Get("key2")
	fmt.Println("res at key2: ", res)
	res, _ = db.Get("key3")
	fmt.Println("res at key3: ", res)

	fmt.Println("reading all")
	iter := db.ReadRange(nil)
	for ; iter.Valid(); iter.Next() {
		fmt.Println("row: ", iter.Key(), " ", iter.Value())
	}
	fmt.Println("iter error: ", iter.Error())
	iter.Release()
}
