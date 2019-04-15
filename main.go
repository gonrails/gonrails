package main

import (
	"fmt"

	_ "kalista/config"
	"kalista/models"
	"kalista/utils/common"
	_ "kalista/utils/sources"
	"time"
)

func main() {
	defer models.Close() // 程序退出之后需要关闭数据库

	models.Tenant{}.Count()

	p := fmt.Println

	t := time.Now()

	p(t.Format(common.DateTimeFormatWithZone))

	p(t.Format(common.DateFormat))

}
