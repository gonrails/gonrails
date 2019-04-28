/*
Author: 曾涛
Desc:   Controllers 的一些通用解析参数方法、辅助方法将会放在这里
Date:   2019-04-28
Email:  zengtao@risewinter.com
*/

package controllers

import (
	"net/url"
)

func QueryToMap(query url.Values) map[string]string {

	result := make(map[string]string)

	for key := range query {
		result[key] = query.Get(key)
	}

	return result
}
