package model

// Model 公共model，大家都有的属性
type Model struct {
	ID         uint32 `gorm:"primary_key" json:"id"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	CreatedOn  string `json:"created_on"`
	ModifiedOn string `json:"modified_on"`
	DeletedOn  uint32 `json:"deleted_on"`
	IsDel      uint8  `json:"is_del"`
}

// Tag 博客标签对应的model
type Tag struct {
	*Model
	Name  string `json:"name"`
	State uint8  `json:"state"`
}

// TableName Tag的数据库中的对应的表名
func (t Tag) TableName() string {
	return "blog_tag"
}

// Article 文章对应的model
type Article struct {
	*Model
	Title         string `json:"title"`
	Desc          string `json:"desc"`
	CoverImageUrl string `json:"cover_image_url"`
	Content       string `json:"content"`
	State         uint8  `json:"state"`
}

// TableName Article的数据库中对应的表名
func (a Article) TableName() string {
	return "blog_article"
}

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
