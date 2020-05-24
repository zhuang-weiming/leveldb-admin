package main

import (
	"fmt"
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

	for i := 0; i < 100; i++ {
		db.Put([]byte(fmt.Sprintf("%s%d", "qyiuqyieiueioquoiueoiquio/qiouoiueo/qio", i)), []byte("valdskjhdkjshkue"), nil)
	}

	<-c
}
