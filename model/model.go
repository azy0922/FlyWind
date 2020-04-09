package model

import (
	"log"
	"strings"

	"github.com/azy0922/flywind/config"
	"github.com/jinzhu/gorm"

	//_ "github.com/jinzhu/gorm/dialects/mysql"
	//_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var (
	DB     *gorm.DB
	dbType = config.String("dbtype")
	dbURL  = config.String("dburl")
)

func init() {
	var err error
	DB, err = gorm.Open(strings.ToLower(dbType), dbURL)
	if err != nil {
		log.Fatalln("connect to database error:", err)
	}
	// 创建表结构
	DB.AutoMigrate(&Problem{})
	DB.AutoMigrate(&Term{})
}
