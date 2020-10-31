package leveldb_admin

import (
	"github.com/siddontang/go/hack"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/util"
	"net/http"
	"time"
)

type apiKeysCountRes struct {
	Count   uint64
	LastKey string
	IsTrue  bool
}

func (l *LevelAdmin) apiKeysCount(writer http.ResponseWriter, request *http.Request) {
	db := request.URL.Query().Get("db")
	if db == "" {
		http.NotFound(writer, request)
		return
	}

	prefix := request.URL.Query().Get("prefix")
	searchText := request.URL.Query().Get("searchText")

	res := &apiKeysCountRes{IsTrue: true, LastKey: searchText}

	timeOut := time.After(time.Second * 1)

	if load, ok := l.dbs.Load(db); ok {
		db := load.(*leveldb.DB)

		iter := db.NewIterator(util.BytesPrefix([]byte(prefix)), nil)
		defer iter.Release()

		if searchText != "" {
			iter.Seek(hack.Slice(searchText))
		}

		for iter.Next() {
			select {
			case <-timeOut:
				res.IsTrue = false
				goto end
			default:
				res.LastKey = string(iter.Key())
				res.Count++
			}
		}
	end:
		err := iter.Error()
		if err != nil {
			l.writeError(writer, err)
			return
		}

		l.writeJson(writer, res)
	} else {
		http.NotFound(writer, request)
	}
}
