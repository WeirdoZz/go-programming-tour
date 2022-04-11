package dao

import "gorm.io/gorm"

type Dao struct {
	engine *gorm.DB
}

// New 新建一个Dao
func New(engine *gorm.DB) *Dao {
	return &Dao{engine: engine}
}
