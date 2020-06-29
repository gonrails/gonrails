/*
Author: 曾涛
Desc:   GET /admin/tenants/:id 的序列化描述
Date:   2019-04-23
Email:  zengtao@risewinter.com
*/

package tenant

//import (
//	"github.com/gonrails/gonrails/models"
//	"github.com/w-zengtao/struct2json"
//)
//
//// AdminTenantIndexSerializer - GET /admin/tenants
//type AdminTenantIndexSerializer struct {
//	*BaseTenantSerializer `json:"base" noroot:"true"`
//}
//
//// Serialize makes AdminTenantIndexSerializer satisfy Serializer interface
//func (s *AdminTenantIndexSerializer) Serialize(v interface{}) map[string]interface{} {
//
//	if obj, ok := v.(*models.Tenant); ok {
//		s = &AdminTenantIndexSerializer{
//			&BaseTenantSerializer{
//				ID:     obj.ID,
//				Name:   obj.Name.String,
//				APIKey: obj.ApiKey,
//			},
//		}
//		ans, _ := struct2json.Convert(s)
//		return ans
//	}
//
//	return map[string]interface{}{"error": "params passed to method Serialize(v interface{}) is not a Tenant"}
//}
