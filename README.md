# leveldb 的web简单管理工具 [嵌入go代码中]

## 使用说明

### 环境变量说明

```
LEVEL_ADMIN_ADDRESS // http服务监听地址 :4333 默认由系统分配端口
LEVEL_ADMIN_DEBUG  // 开启debug true
```

### 使用示例

#### 1. 独立server模式

[LEVEL_ADMIN_ADDRESS=:4333]

```go
package main

import (
    levelAdmin "github.com/qjues/leveldb-admin"
    "github.com/syndtr/goleveldb/leveldb"
)

func main() {
   db, _ := leveldb.OpenFile("/tmp/leveldb", nil)
   // 正常的db操作逻辑...
   // 正常的db操作逻辑...
   
   // 独立端口模式: 
   // 只需要向leveldbAdmin注册db的指针和一个用于区分的描述
   levelAdmin.GetLevelAdmin().Register(db, "description").Start()
}
```

1. 启动后会有一条日志输出[LEVEL_ADMIN_DEBUG=true]时

2020/05/26 00:31:54 leveldb admin server on: http://127.0.0.1:4333/leveldb_admin/static/

#### 2. 共用server模式

[将不处理 LEVEL_ADMIN_ADDRESS 配置]

```go
package main

import (
    levelAdmin "github.com/qjues/leveldb-admin"
    "github.com/syndtr/goleveldb/leveldb"
    "net/http"
)

func main() {
    // 其他server 处理
    http.HandleFunc("/other", func(writer http.ResponseWriter, request *http.Request) {
        writer.Write([]byte("hello other"))
    })
    go http.ListenAndServe(":4333", nil)

    // 正常的db操作逻辑...
    db, _ := leveldb.OpenFile("/tmp/leveldb", nil)
    // 正常的db操作逻辑...

    //levelAdmin 路由全以 leveldb_admin 开头, 可以与其他http server共用server, 需要手动设置 ServerMux
    levelAdmin.GetLevelAdmin().Register(db, "description").SetServerMux(http.DefaultServeMux).Start()
}
```
### 浏览器访问 http://127.0.0.1:4333/leveldb_admin/static/

![demo](https://qjues.github.io/level-db-admin/demo.png)