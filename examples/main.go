package main

import (
	leveldbweb "github.com/Dowte/leveldb-web"
	"github.com/syndtr/goleveldb/leveldb"
	"os"
)

func main() {
	c := make(chan interface{})

	os.MkdirAll("/tmp/leveldb", 0755)
	db, err := leveldb.OpenFile("/tmp/leveldb", nil)
	if err != nil {
		panic(err)
	}

	leveldbweb.Register(db, "temp")

	db.Put([]byte("key"), []byte("value"), nil)

	<-c
}
