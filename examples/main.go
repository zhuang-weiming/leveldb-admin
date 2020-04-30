package main

import (
	"github.com/syndtr/goleveldb/leveldb"
	leveldb_web "leveldb-web"
	"os"
)

func main() {

	c := make(chan interface{})

	os.MkdirAll("/tmp/leveldb", 0755)
	db, err := leveldb.OpenFile("/tmp/leveldb", nil)
	if err != nil {
		panic(err)
	}

	leveldb_web.Register(db, "temp")

	db.Put([]byte("key"), []byte("vale"), nil)

	<-c
}
