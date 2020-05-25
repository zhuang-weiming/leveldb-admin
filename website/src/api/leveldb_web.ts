import http from './http'

export function dbs() {
  return http.get("/leveldb_web/api/dbs")
}

interface KeysOption {
  db: string;
  searchText: string;
  prefix: string;
}

export function keys(option: KeysOption) {
  return http.get("/leveldb_web/api/db/keys", {
    params: option
  })
}

interface KeyInfoOption {
  db: string;
  key: string;
}

export function keyInfo(option: KeyInfoOption) {
  return http.get("/leveldb_web/api/db/key/info", {
    params: option
  })
}

interface KeyDeleteOption {
  db: string;
  key: string;
}

export function keyDelete(option: KeyDeleteOption) {
  return http.post("/leveldb_web/api/db/key/delete", option)
}

interface KeyUpdateOption {
  db: string;
  key: string;
  value: string;
}

export function keyUpdate(option: KeyUpdateOption) {
  return http.post("/leveldb_web/api/db/key/update", option)
}