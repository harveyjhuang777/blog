package storage

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/jwjhuang/blog/service/app/logger"
	"github.com/jwjhuang/blog/service/model"
	"github.com/jwjhuang/blog/service/utils/conf"
)

func NewGORM() *gorm.DB {
	logger.Log().Info("Init GORM")
	connInfo, err := conf.GenDatabaseUrl()
	if err != nil {
		logger.Log().Panic("gen db url failed")
	}

	db, err := gorm.Open(connInfo.Scheme, fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", connInfo.Host, connInfo.Port, connInfo.Username, connInfo.Password, connInfo.DBname))
	if err != nil {
		logger.Log().Panic(err.Error())
	}
	db.DB().SetMaxIdleConns(2)
	db.DB().SetMaxOpenConns(100)

	db.LogMode(true) // for debug

	migrateSchemas(db)

	return db
}

func migrateSchemas(gormdb *gorm.DB) {
	logger.Log().Info("Start migrate schemas")
	db := gormdb
	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Article{})
	db.AutoMigrate(&model.Tag{})
}
