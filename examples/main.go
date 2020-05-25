package main

import (
	"fmt"
	levelAdmin "github.com/Dowte/leveldb-admin"
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

	levelAdmin.Register(db, "temp")

	for i := 0; i < 100; i++ {
		db.Put([]byte(fmt.Sprintf("%s%d", "qyiuqyieiueioquoiueoiquio/qiouoiueo/qio", i)), []byte("valdskjhdkjshkue"), nil)
	}

	<-c
}
