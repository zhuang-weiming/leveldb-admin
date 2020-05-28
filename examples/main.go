package main

import (
	"fmt"
	levelAdmin "github.com/qjues/leveldb-admin"
	"github.com/syndtr/goleveldb/leveldb"
	"net/http"
	"os"
)

func main() {
	c := make(chan interface{})
	http.HandleFunc("/other", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("hello other"))
	})

	go http.ListenAndServe(":4333", nil)

	os.MkdirAll("/tmp/leveldb", 0755)
	db, err := leveldb.OpenFile("/tmp/leveldb", nil)
	if err != nil {
		panic(err)
	}

	levelAdmin.GetLevelAdmin().Register(db, "temp").SetServerMux(http.DefaultServeMux).Start()

	for i := 0; i < 100; i++ {
		db.Put([]byte(fmt.Sprintf("%s%d", "qyiuqyieiueioquoiueoiquio/qiouoiueo/qio", i)), []byte("valdskjhdkjshkue"), nil)
	}

	<-c
}
