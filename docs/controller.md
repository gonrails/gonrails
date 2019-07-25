### [controllers](https://github.com/one-hole/kalista/tree/master/controllers)

  * Package `controllers` stores the implementations to handler request actions
  * We handler single `resource` or a collection of `resources` in __Restful__ world

---

### Examples Explanation

#### 1. `GET /home`

  0. `curl http://localhost:8080/home`

  1. file: `routes/base.go`

  ```go
  root.GET("/home", home.Index) // line 32
  ```

  2. file: `controllers/home/index.go`

  ```go
  package home

  import (
    "github.com/gin-gonic/gin"
  )

  func Index(ctx *gin.Context) {
    ctx.JSON(200, gin.H{
      "message": "Gonrails Starts",
    })
  }

  ```

  3. We route `GET /home` to __function__ `package kalista/controllers/home Index`

#### 2. `GET /books` (Get Collection of resources [books])

  0. `curl http://localhost:8080/books`

  1. file: `routes/base.go`

  ```go
  root.GET("/books", books.Index)
  ```

  2. file: `controllers/books/index`

  ```go
  func Index(ctx *gin.Context) {

  }
  ```

#### 3. `GET /books/:id` (Show single book resource by :id)

  0. `curl http://localhost:8080/books/1`

  1. file: `routes/base.go`

  ```go
  root.GET("/books/:id", books.Show)
  ```

  2. file: `controllers/books/show`

  ```go
  package books

  import (
  	serializer "kalista/serializers"

  	"github.com/gin-gonic/gin"
  	"github.com/one-hole/gonrails/serializers"
  )

  func Show(ctx *gin.Context) {
  	ctx.JSON(200, gin.H{
  		"data": serializers.SingleSerializer(&serializer.BookSerializer{}, nil),
  	})
  }
  ```
  3. file: `serializers/book_serializer.go`

  ```go
  package serializers

  import (
  	"github.com/w-zengtao/struct2json"
  )

  type BookSerializer struct {
  	ID   uint   `json:"id"`
  	Name string `json:"name"`
  }

  func (s *BookSerializer) Serialize(v interface{}) map[string]interface{} {

  	s = &BookSerializer{
  		ID:   1,
  		Name: "Gonrails Demo",
  	}

  	ans, err := struct2json.Convert(s)

  	if nil == err {
  		return ans
  	}

  	return map[string]interface{}{
  		"error": "BookSerializer Serialize Error",
  	}
  }
  ```
  4. Sample Response

  ```json
  {
    "data": {
      "id": 1,
      "name": "Gonrails Demo"
    }
  }
  ```
  5. more infos about package `serializers` click [here](https://github.com/one-hole/gonrails/blob/master/serializers/base.go)

  6. more infos about package `struct2json` click [here](http://github.com/w-zengtao/struct2json)

### Helper Functions
