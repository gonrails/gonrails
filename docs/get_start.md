
### Get Start

#### 0. Install

`go get -u github.com/gonrails/gonrails`

#### 1. Create a project

`gonrails new kalista`

#### 2. Project Overview

* 2.0 Set Configurations

    * Open and edit file __config/config.yml__
    * Configure _Database_
    * Configure _Redis_
    * Configure _RabbitMQ_

* 2.1 Run

    * `GO_ENV=debug go run main.go`

* 2.2 Framework Examples

    * `curl http://localhost:8080/home`
    * `curl http://localhost:8080/admin/home`
    * `curl http://localhost:8080/books/1`
    * more infos at [controller docs](https://github.com/gonrails/gonrails/blob/docs/docs/controller.md)

* 2.3 Directory (including default examples)

    ```
    ├── config                  // Project configurations
    │   ├── app.yml
    │   ├── config.go
    │   └── config.yml
    ├── controllers             // Controller - Handler Request and do Respondse in Restful
    │   ├── base.go
    │   ├── books
    │   │   ├── index.go
    │   │   └── show.go
    │   └── home
    │       └── index.go
    ├── go.mod
    ├── go.sum
    ├── helper
    ├── logs
    ├── main.go
    ├── models                  // Models - Database connections and database operations
    ├── routes                  // Multiplexer - Route request to controller actions
    │   ├── admin
    │   │   └── base.go
    │   └── base.go
    ├── serializers             // Serializers - Do serializer jobs while controller do response
    │   └── book_serializer.go
    └── workers                 // Async jobs
    ```

* 2.4 Dockerfile
    * [example dockerfile](https://github.com/gonrails/gonrails/blob/docs/Dockerfile)

* 2.5 Logs
    * TODO
