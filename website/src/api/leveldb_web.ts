import http from './http'

export function getDbs() {
  return http.get("/leveldb_web/api/dbs")
}

interface GetKeysOption {
  db: string;
  searchText: string;
  prefix: string;
}

export function getKeys(option: GetKeysOption) {
  return http.get("/leveldb_web/api/db/keys", {
    params: option
  })
}

interface GetValueOption {
  db: string;
  key: string;
}

export function getValue(option: GetValueOption) {
  return http.get("/leveldb_web/api/db/value", {
    params: option
  })
}