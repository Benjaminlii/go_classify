package drivers

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"go_classify/biz/config"
	"go_classify/biz/domain/model"
	"log"
)

var (
	DB *gorm.DB
)

func InitMySQL(config *config.BasicConfig) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", config.Mysql.UserName, config.Mysql.PassWorld, config.Mysql.Host, config.Mysql.Port, config.Mysql.DB)
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("[system][initMysql] open mysql error, dsn=%s", dsn)
	}
	db.DB().SetMaxIdleConns(100)
	db.DB().SetMaxOpenConns(1000)

	// 启用Logger，显示详细日志
	db.LogMode(true)

	DB = db
	log.Print("[system][mysql] init mysql driver success!")

	// 数据库建表
	createTable()
}

// createTable 根据模型创建数据库表
func createTable() {
	db := DB
	if !db.HasTable(&model.User{}) {
		db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(&model.User{})
	}
	if !db.HasTable(&model.GarbageDetail{}) {
		db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(&model.GarbageDetail{})
	}
	if !db.HasTable(&model.GarbageType{}) {
		db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(&model.GarbageType{})
	}
	if !db.HasTable(&model.Image{}) {
		db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(&model.Image{})
	}
	if !db.HasTable(&model.ClassifyRecord{}) {
		db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(&model.ClassifyRecord{})
	}
	log.Print("[system][mysql] create table success!")
}
