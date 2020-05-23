import http from './http'

export function getDbs() {
  return http.get("/leveldb_web/api/dbs")
}