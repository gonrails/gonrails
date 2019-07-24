#### [controllers](https://github.com/one-hole/kalista/tree/master/controllers)

  * Package `controllers` stores the implementations to handler request actions
  * We handler single `resource` or a collection of `resources` in __Restful__ world

  
  
#### Examples Explanation

##### `GET /home`

1. `routes/base.go`

```go
root.GET("/home", home.Index) // line 32
```

2. `controllers/home/index.go`

##### `GET /books`

##### `GET /books/:id`