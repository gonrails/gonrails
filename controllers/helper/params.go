package helper

import (
	"net/url"
)

/*
File :  params.go
Author: w-zengtao
Email:  so.zengtao@gmail.com
Desc:   这里提供 Controller 的一些 Helper 方法
*/

/* QueryToMap makes url query params to map[string]string
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
