package leveldb_web

import (
	"encoding/json"
	"fmt"
	"github.com/syndtr/goleveldb/leveldb"
	"log"
	"net"
	"net/http"
	"os"
	"sync"
)

const apiTestUrl = "/leveldb_web/test"
const staticPrefix = "/leveldb_web/static/"

const (
	apiPrefix = "/leveldb_web/api"
	apiDbs    = apiPrefix + "/dbs"
	apiKeys   = apiPrefix + "/db/keys"
	apiValue  = apiPrefix + "/db/value"
)

type LevelWeb struct {
	dbs     sync.Map
	address string
	debug   bool
	mux     *http.ServeMux
}

var levelWeb = &LevelWeb{}

// Register after init
func Register(db *leveldb.DB, key string) {
	levelWeb.logInfo(fmt.Sprintf("add db register: %s, %p", key, db))

	levelWeb.dbs.Store(key, db)
}

func init() {
	if envAddr := os.Getenv("LEVEL_WEB_ADDRESS"); envAddr != "" {
		levelWeb.address = envAddr
	}

	if envAddr := os.Getenv("LEVEL_WEB_DEBUG"); envAddr == "true" {
		levelWeb.debug = true
	}

	go levelWeb.startServer()
}

func (l *LevelWeb) apiHelloWord(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("hello world"))
}

func (l *LevelWeb) startServer() error {
	listen, err := net.Listen("tcp", l.address)

	if err != nil {
		return err
	}
	l.mux = http.NewServeMux()

	l.startStatic(staticPrefix)

	l.mux.HandleFunc(apiTestUrl, l.apiHelloWord)
	l.mux.HandleFunc(apiDbs, l.apiDBs)
	l.mux.HandleFunc(apiKeys, l.apiKeys)
	l.mux.HandleFunc(apiValue, l.apiValue)

	port := listen.Addr().(*net.TCPAddr).Port

	server := http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: l.mux,
	}

	log.Printf("leveldb web server on: http://%s:%d", "127.0.0.1", listen.Addr().(*net.TCPAddr).Port)

	return server.Serve(listen)
}

func (l *LevelWeb) writeError(writer http.ResponseWriter, err error) {
	writer.Header().Add("Content-Type", "application/json")

	_, _ = writer.Write([]byte(fmt.Sprintf("{\"error:\" %s}", err.Error())))
}

func (l *LevelWeb) writeJson(writer http.ResponseWriter, v interface{}) {
	marshal, err := json.Marshal(v)
	if err != nil {
		l.writeError(writer, err)
	} else {
		writer.Header().Add("Content-Type", "application/json")

		_, _ = writer.Write(marshal)
	}
}

func (l *LevelWeb) logInfo(str string) {
	if l.debug {
		log.Println(str)
	}
}

func (l *LevelWeb) logInfoWithFunc(c func()) {
	if l.debug {
		c()
	}
}
