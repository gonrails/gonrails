### Model

* package `models` is the place where communication between database and business happens. And it bridges golang structs to database tables.


#### Helper Functions

  * `func GenerateToken() string` - default `[]byte` length is 16

#### MySQL Helper

  * `func Open(host, port, username, password, name string) *gorm.DB`
  * `func ConfigureMySQL(db *gorm.DB, idles, connections int)`



More infomations of `GORM` click [here](https://gorm.io)
