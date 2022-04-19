package model

import (
	"blog-service/global"
	"blog-service/pkg/setting"
	"fmt"
	"gorm.io/driver/mysql"
	_ "gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"time"
)

// Model 公共model，大家都有的属性
type Model struct {
	ID         uint32 `gorm:"primary_key" json:"id"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	CreatedOn  uint32 `json:"created_on"`
	ModifiedOn uint32 `json:"modified_on"`
	DeletedOn  uint32 `json:"deleted_on"`
	IsDel      uint8  `json:"is_del"`
}

// NewDBEngine 创建一个新的DBEngine
func NewDBEngine(databaseSetting *setting.DatabaseSettingS) (*gorm.DB, error) {
	s := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=Local",
		databaseSetting.Username,
		databaseSetting.Password,
		databaseSetting.Host,
		databaseSetting.DBName,
		databaseSetting.Charset,
		databaseSetting.ParseTime)
	db, err := gorm.Open(mysql.Open(s), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   databaseSetting.TablePrefix,
			SingularTable: true,
		},
	})
	if err != nil {
		return nil, err
	}
	if global.ServerSetting.RunMode == "debug" {
		db.Logger = logger.Default.LogMode(logger.Info)
	}
	sqlDB, err := db.DB()

	sqlDB.SetMaxIdleConns(databaseSetting.MaxIdleConns)
	sqlDB.SetMaxOpenConns(databaseSetting.MaxOpenConns)

	// 注册回调函数
	db.Callback().Create().Replace("gorm:update_time_stamp", updateTimeStampForCreateCallback)
	db.Callback().Update().Replace("gorm:update_time_stamp", updateTimeStampForUpdateCallback)
	db.Callback().Delete().Replace("gorm:delete", deleteCallback)
	return db, nil
}

// updateTimeStampForCreateCallback 更新创建新标签时间的回调函数
func updateTimeStampForCreateCallback(db *gorm.DB) {
	db.Statement.SetColumn("CreatedOn", time.Now().Unix(), true)
}

// updateTimeStampForUpdateCallback 更新修改标签的时间的回调函数
func updateTimeStampForUpdateCallback(db *gorm.DB) {
	db.Statement.SetColumn("ModifiedOn", time.Now().Unix(), true)
}

// deleteCallback 删除标签时对删除时间的回调函数
func deleteCallback(db *gorm.DB) {
	if db.Error == nil {
		if db.Statement.Schema != nil {
			db.Statement.SQL.Grow(100)

			deleteField := db.Statement.Schema.LookUpField("DeletedOn")
			if !db.Statement.Unscoped && deleteField != nil {
				// 实现软删除
				if db.Statement.SQL.String() == "" {
					// 软删除
					nowTime := time.Now().Unix()
					db.Statement.AddClause(
						clause.Set{
							{Column: clause.Column{Name: deleteField.DBName},
								Value: nowTime,
							}},
					)

					db.Statement.AddClauseIfNotExists(clause.Update{})
					db.Statement.Build("UPDATE", "SET", "WHERE")
				}
			} else {
				if db.Statement.SQL.String() == "" {
					db.Statement.AddClauseIfNotExists(clause.Delete{})
					db.Statement.AddClauseIfNotExists(clause.From{})
					db.Statement.Build("DELETE", "FROM", "WHERE")
				}
			}

			// 必须要有WHERE
			if _, ok := db.Statement.Clauses["WHERE"]; !db.AllowGlobalUpdate && !ok {
				db.AddError(gorm.ErrMissingWhereClause)
				return
			}

			if !db.DryRun && db.Error == nil {
				result, err := db.Statement.ConnPool.ExecContext(db.Statement.Context, db.Statement.SQL.String(), db.Statement.Vars...)
				if err == nil {
					db.RowsAffected, _ = result.RowsAffected()
				} else {
					db.AddError(err)
				}
			}
		}
	}
}
