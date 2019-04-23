/*
Author: 曾涛
Desc:   GET /admin/tenants/:id 的序列化描述
Date:   2019-04-23
Email:  zengtao@risewinter.com
*/

package tenant

import (
	"kalista/models"
	"log"
)

type AdminTenantShowSerializer struct {
	baseTenantSerializer
}

func (s *AdminTenantShowSerializer) Serialize(v interface{}) map[string]interface{} {

	if s, ok := v.(*models.Tenant); ok {
		log.Println(s)
		return map[string]interface{}{
			"id": "right",
		}
	} else {
		return map[string]interface{}{
			"id": "abc",
		}
	}
}
