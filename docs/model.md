### Model

* package `models` is the place we communicate between database and business. And it provide an abstraction for golang struct to database table.


#### Helper Functions

  * `func GenerateToken() string` - default `[]byte` length is 16

#### MySQL Helper

  * `func Open(host, port, username, password, name string) *gorm.DB`
  * `func ConfigureMySQL(db *gorm.DB, idles, connections int)`



More infomations of `GORM` click [here](https://gorm.io)
