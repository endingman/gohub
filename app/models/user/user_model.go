// Package user 存放用户 Model 相关逻辑
package user

import (
	"gohub/app/models"
)

// User 用户模型

type User struct {
	models.BaseModel
	// 不希望将敏感信息输出给用户，所以这里 Email 、Phone 、Password 后面设置了 json:"-" ，
	// 这是在指示 JSON 解析器忽略字段 。
	// 后面接口返回用户数据时候，这三个字段都会被隐藏
	Name     string `json:"name,omitempty"`
	Email    string `json:"-"`
	Phone    string `json:"-"`
	Password string `json:"-"`

	models.CommonTimestampsField
}
