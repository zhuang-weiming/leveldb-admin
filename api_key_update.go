package leveldb_admin

import (
	"encoding/json"
	"github.com/siddontang/go/hack"
	"github.com/syndtr/goleveldb/leveldb"
	"net/http"
)

type updateRes struct {
	Success bool
}

type updateReq struct {
	DB    string
	Key   string
	Value string
}

func (l *LevelAdmin) apiKeyUpdate(writer http.ResponseWriter, request *http.Request) {
	reqData := &updateReq{}
	err := json.NewDecoder(request.Body).Decode(&reqData)
	if err != nil {
		l.writeError(writer, err)

		return
	}

	if reqData.DB == "" || reqData.Key == "" {
		http.NotFound(writer, request)
		return
	}

	if load, ok := l.dbs.Load(reqData.DB); ok {
		db := load.(*leveldb.DB)
		if has, err := db.Has(hack.Slice(reqData.Key), nil); has && err == nil {
			err := db.Put(hack.Slice(reqData.Key), []byte(reqData.Value), nil)
			if err != nil {
				l.writeError(writer, err)

				return
			}
		} else {
			http.NotFound(writer, request)
			return
		}

		l.writeJson(writer, &updateRes{Success: true})
	} else {
		http.NotFound(writer, request)
	}
}
