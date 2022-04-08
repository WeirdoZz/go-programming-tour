package model

// ArticleTag 文章和标签之间的对应关系表
type ArticleTag struct {
	*Model
	TagID     uint32 `json:"tag_id"`
	ArticleID uint32 `json:"article_id"`
}

// TableName Article和Tag关系的数据库中对应的表名
func (a ArticleTag) TableName() string {
	return "blog_article_tag"
}
