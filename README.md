[Gonrails](https://github.com/one-hole/gonrails) 是基于 [Gin](https://github.com/gin-gonic/gin) 改造的后端 MVC 开发框架

[![Go Report Card](https://goreportcard.com/badge/github.com/one-hole/kalista)](https://goreportcard.com/report/github.com/one-hole/kalista)
[![MIT license](http://img.shields.io/badge/license-MIT-brightgreen.svg)](http://opensource.org/licenses/MIT)
--------

* QQ 群： 594171535
* 接头暗号： Heil Hydra

#### 下一步即将要集成的事情

  * 将 gonrails-cli 项目集成到本项目里面来
  * 更好的错误输出（也就是 API 交互调用者的错误代码）
  * 分布式，服务的注册、发现，请求分发（master - slave）。
  * 该项目最终的形态应该是一个插件类型的项目、提供工具函数、提供通用方法。辅助 Gonrails-cli 项目

  比如:

  import "github.com/one-hole/gonrails/cluster" 这个包、就能启用集群（主从）的组件


#### Cli 工具

  * [gonrails-cli](https://github.com/one-hole/gonrails-cli)

### 如何启动项目:

  #### Development

  ```shell
  GO_ENV=debug go run main.go
  ```
  #### Production

  ```shell
  GO_ENV=release ./kalista
  ```


### 启动之前你需要做些什么：

  * 编辑 config/config.yml 确保可用
  * 确认 数据库、Redis、消息中间件等服务可用（如果不想用、可以在[相应的地方](https://github.com/one-hole/kalista/blob/master/utils/sources/base.go)注释掉）

    ```go
    package sources

    import (
      "kalista/utils/sources/rabbits"
    )

    func init() {
      go rabbits.RunRabbit() // 注释掉这行、就不会去连接 RabbitMQ
    }
    ```
  * 一些日志的配置（建议自定义日志输出、包括 [Gorm](https://github.com/jinzhu/gorm) & [Gin](https://github.com/gin-gonic/gin)）
---

项目目录结构如下（部分叶子目录结构包含使用案例）

  ```
  ├── config
  ├── console
  ├── controllers
  │   ├── admin
  │   │   ├── reports
  │   │   │   ├── revenues
  │   │   │   ├── statistics
  │   │   │   └── topics
  │   │   └── tenants
  │   ├── outer
  │   └── websocket
  ├── models
  ├── routers
  │   ├── admin
  │   │   └── reports
  │   └── outer
  ├── serializers
  │   ├── report
  │   └── tenant
  ├── service
  ├── utils
  │   ├── common
  │   ├── deprecated
  │   └── sources
  │       └── rabbits
  └── workers
  ```


#### [Config](https://github.com/one-hole/kalista/tree/master/config)

  config 目录（包）主要用来存放配置文件、我遵循 Gin 的 Mode 设置了 `debug` 和 `release` 两个模式

  sample 如下

  ```yml
  debug:
    MySQL:
      host: localhost
      port: 3306
      name: rw-datacenter-dev-golang
      username: root
      password: mypassword
      connections: 60
      idles: 10
    Redis:
      host: 127.0.0.1
      port: 6379
    RabbitMQ:
      host: 127.0.0.1
      port: 5672
      user: guest
      password: guest

  release:
  ```

  个人建议：加入日志的配置在这里会更加合理一点、在这个事情上我有点懒了

#### [Models](https://github.com/one-hole/kalista/tree/master/models)

  __我们是支持多 数据库 的__
  * [mysql](https://github.com/one-hole/kalista/blob/master/models/mysql.go#L14) 这里多申明一个变量、保存创建连接即可

    ```go
    var (
      db *gorm.DB
      anotherDB *gorm.DB //另外一个DB，然后你需要在下面 init() 中， open() 一下
    )
    ```
  * models 目录（包）这里主要保存的是数据库的映射关系（ORM）和 数据连接对象
  * _scopes.go 结尾的文件主要是定义一下查询的过滤条件、定义通用的 `FilterParams` 方法, 使得在 `Controller` 里面的查询行为一致、并且查询行为可以由前端来控制参数

    [应用示例](https://github.com/one-hole/kalista/blob/master/controllers/admin/reports/statistics/index.go)

    ```go
    package statistics

    import (
      "kalista/controllers"
      "kalista/models"
      "kalista/serializers"
      "kalista/serializers/report"
      "net/http"

      "github.com/gin-gonic/gin"
    )

    func Index(ctx *gin.Context) {
      records, record := []*models.StatisticReports{}, &models.StatisticReports{}
      query := ctx.Request.URL.Query()
      record.FilterParams(controllers.QueryToMap(query)).Find(&records)

      resp, err := serializers.CollectionSerializer(&report.AdminStatisticReportIndexSerializer{}, records)

      if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{
          "data": err,
        })
      }

      ctx.JSON(http.StatusOK, gin.H{
        "data": resp,
      })
    }
    ```

#### [Routers](https://github.com/one-hole/kalista/tree/master/routers)
  * 配置、启动 Gin & HTTP 相关参数
  * 路由、以及相关函数映射（也就是路由映射到资源的 Index | Create | Update | Delete | Show 之类的 Rest 行为上）

    示例 - [https://github.com/one-hole/kalista/blob/master/routers/admin/base.go](https://github.com/one-hole/kalista/blob/master/routers/admin/base.go)
    ```go
    package admin

    import (
      "kalista/controllers/admin/tenants"
      "kalista/routers/admin/reports"

      "github.com/gin-gonic/gin"
    )

    func ApplyRoutes(router *gin.RouterGroup) {
      adminGroup := router.Group("/admin")
      {
        reports.ApplyRoutes(adminGroup)
        //下面将 GET /admin/tenants 路由映射到了 kalista/controllers/admin/tenants 这个包的 Index 函数上
        adminGroup.GET("/tenants", tenants.Index)

        //下面将 GET /admin/tenants/:id 路由映射到了 kalista/controllers/admin/tenants 这个包的 Show 函数上
        adminGroup.GET("/tenants/:id", tenants.Show)
        adminGroup.POST("/tenants", tenants.Create)
      }
    }
    ```

#### [Controllers](https://github.com/one-hole/kalista/tree/master/controllers)

  * Controllers 目录（包）主要是 APi 的具体实现、承接路由的落地行为（Actions)
  * 在 Restful 的世界里，我们对单个资源进行操作 (Show, Update, Post, Delete)、或者对多个资源集合操作 (Index)

    下面给出一个目录示例:

    ```
    ├── admin
    │   ├── base.go
    │   ├── reports
    │   │   ├── base.go
    │   │   ├── revenues
    │   │       ├── index.go
    │   │       └── show.go
    │   └── tenants
    │       ├── base.go
    │       ├── create.go
    │       ├── index.go
    │       └── show.go
    ├── base.go
    ├── helper.go
    ├── outer
    │   └── base.go
    └── websocket
        ├── base.go
        └── readme.md
    ```
  * [helper.go](https://github.com/one-hole/kalista/blob/master/controllers/helper.go) 这里会存放一些通用的方法、比如获取当前的分页、比如把 Query 转化成 Map 的一些通用方法

#### [Serializers](https://github.com/one-hole/kalista/tree/master/serializers) 这里也就是我们的 View 层

  * Serializers 目录（包）主要存储我们即将结构化输出的序列化的结构体
  * 定义了两个方法分别用来序列化单个对象 和 序列化集合对象（这两个方法无需改动）

    ```go
    func SingleSerializer(s Serializer, v interface{}) map[string]interface{} {
      return s.Serialize(v)
    }
    ```

    ```go
    func CollectionSerializer(s Serializer, vs interface{}) ([]map[string]interface{}, error) {

      ansAry := []map[string]interface{}{}

      if reflect.Slice != reflect.TypeOf(vs).Kind() {
        return []map[string]interface{}{}, errors.New("Data must be slice type")
      }

      value := reflect.ValueOf(vs)

      for i := 0; i < value.Len(); i++ {
        ansAry = append(ansAry, s.Serialize(value.Index(i).Interface()))
      }

      return ansAry, nil
    }
    ```
  * 每一个用来序列化的结构体定义都需要实现这个方法

    ```go
    type Serializer interface {
      Serialize(v interface{}) map[string]interface{}
    }
    ```
  * 在这里我提供了一个组件 [Struct2Json](https://github.com/w-zengtao/struct2json) 来帮助我们直接把嵌套的对象 转化成 map[string]interface{}

    使用案例 - [statistic_index_serializer.go](https://github.com/one-hole/kalista/blob/master/serializers/report/statistic_index_serializer.go)


#### Docker

```bash
docker build kalista:1.0 .

docker run -e GO_ENV=debug -v "/$(pwd)/config/config.yml:/root/config/config.yml"  kalista:1.0
docker run -e GO_ENV=release -v "/$(pwd)/config/config.yml:/root/config/config.yml"  kalista:1.0

实时日志
docker logs -f your-container-id
```
