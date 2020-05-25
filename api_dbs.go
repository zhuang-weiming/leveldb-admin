package leveldb_admin

import (
	"fmt"
	"net/http"
)

func (l *LevelAdmin) apiDBs(writer http.ResponseWriter, request *http.Request) {
	var dbsMap []string

	l.dbs.Range(func(key, value interface{}) bool {
		dbsMap = append(dbsMap, fmt.Sprintf("%v", key))

		return true
	})

	l.writeJson(writer, dbsMap)
}
