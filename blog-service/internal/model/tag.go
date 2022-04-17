package model

import (
	"blog-service/pkg/app"
	"gorm.io/gorm"
)

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

// Count 计算同名标签个数
func (t Tag) Count(db *gorm.DB) (int64, error) {
	var count int64
	if t.Name != "" {
		db = db.Where("name=?", t.Name)
	}
	db = db.Where("state=?", t.State)
	err := db.Model(&t).Where("is_del=?", 0).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

// List 列出数据库中同名标签
func (t Tag) List(db *gorm.DB, pageOffset, pageSize int) ([]*Tag, error) {
	var tags []*Tag
	var err error

	if pageOffset >= 0 && pageSize > 0 {
		db = db.Offset(pageOffset).Limit(pageSize)
	}

	if t.Name != "" {
		db = db.Where("name=?")
	}
	db = db.Where("state=?", t.State)
	if err = db.Where("is_del=?", 0).Find(&tags).Error; err != nil {
		return nil, err
	}
	return tags, nil
}

// Create 数据库中创建Tag
func (t Tag) Create(db *gorm.DB) error {
	return db.Create(&t).Error
}

// Update 修改数据库中该标签内容
func (t Tag) Update(db *gorm.DB, values interface{}) error {
	err := db.Model(t).Where("id=? AND is_del=?", t.ID, 0).Updates(values).Error
	if err != nil {
		return err
	}
	return nil
}

// Delete 数据库中删除该标签
func (t Tag) Delete(db *gorm.DB) error {
	return db.Where("id=? AND is_del=?", t.ID, 0).Delete(&t).Error
}
