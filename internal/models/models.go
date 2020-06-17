// Copyright © 2020. Drew Lee. All rights reserved.

package models

import (
	"KarlMalone/config"
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

type Model struct {
	ID        int64      `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at" sql:"index"`
}

// init mysql connection by gorm
func init() {
	var err error
	dbType := config.Database.Type
	dbName := config.Database.Name
	user := config.Database.User
	password := config.Database.Password
	host := config.Database.Host
	tablePrefix := config.Database.TablePrefix

	dbArgs := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user,
		password,
		host,
		dbName)
	db, err = gorm.Open(dbType, dbArgs)
	fmt.Printf("%s connect args: %s", dbType, dbArgs)
	if err != nil {
		panic(fmt.Errorf("database init error: %v", err))
	}

	if err := db.DB().Ping(); err != nil {
		panic(fmt.Errorf("dao Ping error: %v", err))
	}

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return tablePrefix + defaultTableName
	}

	db.Callback().Create().Replace("gorm:update_time_stamp", updateTimeStampForCreateCallback)
	db.Callback().Update().Replace("gorm:update_time_stamp", updateTimeStampForUpdateCallback)
	db.Callback().Delete().Replace("gorm:delete", deleteCallback)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	// !! add db logger to check db i/o health
	fmt.Println("database init success.")
}

// CloseDB closes database connection (unnecessary)
func CloseDB() {
	defer db.Close()
}

// updateTimeStampForCreateCallback will set `CreatedAt`, `UpdatedAt` when creating
func updateTimeStampForCreateCallback(scope *gorm.Scope) {
	if !scope.HasError() {
		now := time.Now()

		if createdAtField, ok := scope.FieldByName("CreatedAt"); ok {
			if createdAtField.IsBlank {
				_ = createdAtField.Set(now)
			}
		}

		if updatedAtField, ok := scope.FieldByName("UpdatedAt"); ok {
			if updatedAtField.IsBlank {
				_ = updatedAtField.Set(now)
			}
		}
	}
}

// updateTimeStampForUpdateCallback will set `ModifiedOn` when updating
func updateTimeStampForUpdateCallback(scope *gorm.Scope) {
	if _, ok := scope.Get("gorm:update_column"); !ok {
		_ = scope.SetColumn("UpdatedAt", time.Now())
	}
}

// deleteCallback will set `DeletedAt` where deleting
func deleteCallback(scope *gorm.Scope) {
	if !scope.HasError() {
		var extraOption string
		if str, ok := scope.Get("gorm:delete_option"); ok {
			extraOption = fmt.Sprint(str)
		}

		deletedAtField, hasDeletedAtField := scope.FieldByName("DeletedAt")

		if !scope.Search.Unscoped && hasDeletedAtField {
			scope.Raw(fmt.Sprintf(
				"UPDATE %v SET %v=%v%v%v",
				scope.QuotedTableName(),
				scope.Quote(deletedAtField.DBName),
				scope.AddToVars(time.Now()),
				addExtraSpaceIfExist(scope.CombinedConditionSql()),
				addExtraSpaceIfExist(extraOption),
			)).Exec()
		} else {
			scope.Raw(fmt.Sprintf(
				"DELETE FROM %v%v%v",
				scope.QuotedTableName(),
				addExtraSpaceIfExist(scope.CombinedConditionSql()),
				addExtraSpaceIfExist(extraOption),
			)).Exec()
		}
	}
}

// addExtraSpaceIfExist adds a separator
func addExtraSpaceIfExist(str string) string {
	if str != "" {
		return " " + str
	}
	return ""
}
