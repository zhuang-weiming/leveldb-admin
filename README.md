# leveldb 的web简单管理工具 [嵌入go代码中]

## 使用示例

### 环境变量说明

```
LEVEL_ADMIN_ADDRESS // http服务监听地址 :4333 默认由系统分配端口
LEVEL_ADMIN_DEBUG  // 开启debug true
```

### 代码中

```go
package main

import (
    levelAdmin "github.com/Dowte/leveldb-admin"
    "github.com/syndtr/goleveldb/leveldb"
)

func main() {
   db, _ := leveldb.OpenFile("/tmp/leveldb", nil)
   // 正常的db操作逻辑...
   // 正常的db操作逻辑...

   // 只需要向leveldbAdmin注册db的指针和一个用于区分的描述
   levelAdmin.GetLevelAdmin().Register(db, "description")
}
```
1. 启动后会有一条日志输出[LEVEL_ADMIN_DEBUG=true]时

2020/05/26 00:31:54 leveldb admin server on: http://127.0.0.1:4333/leveldb_admin/static/

### 浏览器访问 http://127.0.0.1:4333/leveldb_admin/static/

![demo](https://raw.githubusercontent.com/Dowte/imgs/master/1057DA66-A4FC-42EE-A3EE-F3FD910DA073.png)