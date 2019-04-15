package main

import (
	_ "kalista/config"
	"kalista/models"
)

func main() {
	defer models.Close() // 程序退出之后需要关闭数据库

	models.Tenant{}.Count()
}
