package model

import "blog-service/pkg/app"

// Tag 博客标签对应的model
type Tag struct {
	*Model
	Name  string `json:"name"`
	State uint8  `json:"state"`
}

type TagSwagger struct {
	List  []*Tag
	Pager *app.Pager
}

// TableName Tag的数据库中的对应的表名
func (t Tag) TableName() string {
	return "blog_tag"
}
