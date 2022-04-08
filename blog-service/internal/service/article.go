package service

type GetArticleRequest struct {
	ID uint32 `form:"id" binding:"required,gte=1"`
}
type ListArticleRequest struct {
	TagID uint32 `form:"tag_id" binding:"required,gte=1"`
	State uint8  `form:"state"`
}
type CreateArticleRequest struct {
	TagID         uint8  `form:"tag_id" binding:"required,gte=1"`
	Title         string `form:"title" binding:"required,max=100"`
	Desc          string `form:"desc" binding:"required,max=255"`
	CoverImageUrl string `form:"cover_image_url" binding:"required,url"`
	Content       string `form:"content" binding:"required,min=2,max=4294967295"`
	CreatedBy     string `form:"created_by" binding:"required,min=3,max=100"`
	State         uint8  `form:"state" binding:"required,oneof=0 1"`
}
type UpdateArticleRequest struct {
	ID            uint32 `form:"id" binding:"required,gte=1"`
	TagID         uint8  `form:"tag_id" binding:"gte=1"`
	Title         string `form:"title" binding:"max=100"`
	Desc          string `form:"desc" binding:"max=255"`
	CoverImageUrl string `form:"cover_image_url" binding:"url"`
	Content       string `form:"content" binding:"min=2,max=4294967295"`
	ModifiedBy    string `form:"modified_by" binding:"required,min=3,max=100"`
	State         uint8  `form:"state" binding:"required,oneof=0 1"`
}
type DeleteArticleRequest struct {
	ID uint32 `form:"id" binding:"required,gte=1"`
}
