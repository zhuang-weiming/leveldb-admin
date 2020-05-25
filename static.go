package leveldb_admin

import (
	"github.com/GeertJohan/go.rice"
	"net/http"
)

//must be run `go get github.com/GeertJohan/go.rice/rice` first
//go:generate rice embed-go

func (l *LevelAdmin) startStatic(prefix string) {
	l.mux.Handle(prefix, http.StripPrefix(prefix, http.FileServer(rice.MustFindBox("static").HTTPBox())))
}
