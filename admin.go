package leveldb_admin

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

const apiTestUrl = "/leveldb_admin/test"
const staticPrefix = "/leveldb_admin/static/"

const (
	apiPrefix    = "/leveldb_admin/api"
	apiDbs       = apiPrefix + "/dbs"
	apiKeys      = apiPrefix + "/db/keys"
	apiKeyInfo   = apiPrefix + "/db/key/info"
	apiKeyDelete = apiPrefix + "/db/key/delete"
	apiKeyUpdate = apiPrefix + "/db/key/update"
)

type LevelAdmin struct {
	dbs     sync.Map
	address string
	debug   bool
	mux     *http.ServeMux
}

var levelAdmin *LevelAdmin
var once sync.Once

func GetLevelAdmin() *LevelAdmin {
	if levelAdmin == nil {
		once.Do(func() {
			levelAdmin = &LevelAdmin{}
			levelAdmin.loadEnv()
		})
	}

	return levelAdmin
}

// Register after init
func (l *LevelAdmin) Register(db *leveldb.DB, key string) *LevelAdmin {
	levelAdmin.logInfo(fmt.Sprintf("add db register: %s, %p", key, db))

	levelAdmin.dbs.Store(key, db)

	return l
}

func (l *LevelAdmin) SetServerMux(mux *http.ServeMux) *LevelAdmin {
	if l.mux == nil {
		l.mux = mux
	}

	return l
}

func (l *LevelAdmin) loadEnv() {
	if envAddr := os.Getenv("LEVEL_ADMIN_ADDRESS"); envAddr != "" {
		l.address = envAddr
	}

	if debug := os.Getenv("LEVEL_ADMIN_DEBUG"); debug == "true" {
		l.debug = true
	}
}

func (l *LevelAdmin) apiHelloWord(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("hello world"))
}

func (l *LevelAdmin) initServerMux() error {
	if l.mux == nil {
		listen, err := net.Listen("tcp", l.address)

		if err != nil {
			l.logInfo(fmt.Sprintf("listen %s error: %v", l.address, err))
			return err
		}
		l.mux = http.NewServeMux()

		port := listen.Addr().(*net.TCPAddr).Port

		server := http.Server{
			Addr:    fmt.Sprintf(":%d", port),
			Handler: l.mux,
		}

		l.logInfo(fmt.Sprintf("leveldb admin server on: http://%s:%d/leveldb_admin/static/", "127.0.0.1", listen.Addr().(*net.TCPAddr).Port))

		go func() {
			err = server.Serve(listen)
			if err != nil {
				l.logInfo(fmt.Sprintf("server on %s error: %v", l.address, err))
			}
		}()
	} else {
		l.logInfo("leveldb admin server on given mux")
	}

	return nil
}

func (l *LevelAdmin) Start() error {
	err := l.initServerMux()

	if err != nil {
		l.logInfo(fmt.Sprintf("init server mux: %v", err))
		return err
	}

	l.startStatic(staticPrefix)

	l.mux.HandleFunc(apiTestUrl, l.apiHelloWord)
	l.mux.HandleFunc(apiDbs, l.apiDBs)
	l.mux.HandleFunc(apiKeys, l.apiKeys)
	l.mux.HandleFunc(apiKeyInfo, l.apiKeyInfo)
	l.mux.HandleFunc(apiKeyDelete, l.apiKeyDelete)
	l.mux.HandleFunc(apiKeyUpdate, l.apiKeyUpdate)

	l.logInfo("leveldb admin server started.")

	return nil
}

func (l *LevelAdmin) writeError(writer http.ResponseWriter, err error) {
	writer.Header().Add("Content-Type", "application/json")

	_, _ = writer.Write([]byte(fmt.Sprintf("{\"error:\" %s}", err.Error())))
}

func (l *LevelAdmin) writeJson(writer http.ResponseWriter, v interface{}) {
	marshal, err := json.Marshal(v)
	if err != nil {
		l.writeError(writer, err)
	} else {
		writer.Header().Add("Content-Type", "application/json")

		_, _ = writer.Write(marshal)
	}
}

func (l *LevelAdmin) logInfo(str string) {
	if l.debug {
		log.Println(str)
	}
}

func (l *LevelAdmin) logInfoWithFunc(c func()) {
	if l.debug {
		c()
	}
}
