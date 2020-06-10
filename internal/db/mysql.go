// Copyright © 2020. Drew Lee. All rights reserved.

package db

import (
	"KarlMalone/config"
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var Orm *gorm.DB

// init mysql connection by gorm
func init() {
	var (
		err                                               error
		dbType, dbName, user, password, host, tablePrefix string
	)
	dbType = config.Database.Type
	dbName = config.Database.Name
	user = config.Database.User
	password = config.Database.Password
	host = config.Database.Host
	tablePrefix = config.Database.TablePrefix

	dbArgs := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user,
		password,
		host,
		dbName)
	Orm, err = gorm.Open(dbType, dbArgs)
	fmt.Printf("%s connect args: %s", dbType, dbArgs)
	if err != nil {
		panic(fmt.Errorf("database init error: %v", err))
	}
	if err := Orm.DB().Ping(); err != nil {
		panic(fmt.Errorf("dao Ping error: %v", err))
	}

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return tablePrefix + defaultTableName
	}
	// !! add db logger to check db i/o health
	fmt.Println("database init success.")
	//defer Orm.Close()
}
