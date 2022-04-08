package model

import "blog-service/pkg/app"

// Article 文章对应的model
type Article struct {
	*Model
	Title         string `json:"title"`
	Desc          string `json:"desc"`
	CoverImageUrl string `json:"cover_image_url"`
	Content       string `json:"content"`
	State         uint8  `json:"state"`
}

type ArticleSwagger struct {
	List  []*Article
	Pager *app.Pager
}

// TableName Article的数据库中对应的表名
func (a Article) TableName() string {
	return "blog_article"
}
