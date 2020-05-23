package leveldb_web

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (l *LevelWeb) apiDBs(writer http.ResponseWriter, request *http.Request) {
	var dbsMap []string

	l.dbs.Range(func(key, value interface{}) bool {
		dbsMap = append(dbsMap, fmt.Sprintf("%v", key))

		return true
	})

	marshal, err := json.Marshal(dbsMap)
	if err != nil {
		writeError(writer, err)
	} else {
		writer.Header().Add("Content-Type", "application/json")
		writer.Write(marshal)
	}
}
