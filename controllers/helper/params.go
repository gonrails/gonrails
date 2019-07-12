package helper

import (
	"net/url"

	"github.com/gin-gonic/gin"
)

/*
File :  params.go
Author: w-zengtao
Email:  so.zengtao@gmail.com
Desc:   这里提供 Controller 的一些 Helper 方法
*/

/*QueryToMap makes url query params to map[string]string
eg:

url: https://yourdomain?a=a&b=b
returns:

map[string]string{
		"a": "a",
		"b": "b",
}
*/
func QueryToMap(query url.Values) map[string]string {
	return queryToMap(query)
}

func queryToMap(query url.Values) map[string]string {
	result := make(map[string]string)

	for key := range query {
		result[key] = query.Get(key)
	}

	return result
}

/*Params returns the request's all params as map[string]interface{}
1. Query Params	(All String)
2. Body Params
3. Form Params
4. URL Params
*/
func Params(ctx *gin.Context) map[string]interface{} {
	return params(ctx)
}

func params(ctx *gin.Context) map[string]interface{} {
	return nil
}
\
