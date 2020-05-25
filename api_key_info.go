package leveldb_web

import (
	"github.com/siddontang/go/hack"
	"github.com/syndtr/goleveldb/leveldb"
	"net/http"
)

type keyInfoRes struct {
	Value interface{}
}

func (l *LevelWeb) apiKeyInfo(writer http.ResponseWriter, request *http.Request) {
	db := request.URL.Query().Get("db")
	key := request.URL.Query().Get("key")
	if db == "" || key == "" {
		http.NotFound(writer, request)
		return
	}

	if load, ok := l.dbs.Load(db); ok {
		db := load.(*leveldb.DB)
		value, err := db.Get(hack.Slice(key), nil)

		if err != nil {
			l.writeError(writer, err)
			return
		}

		l.writeJson(writer, &keyInfoRes{Value: string(value)})
	} else {
		http.NotFound(writer, request)
	}
}
