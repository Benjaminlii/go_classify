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

func createTable() {
	db := DB

	if !db.HasTable(&model.Administrator{}) {
		db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(&model.Administrator{})
		log.Print("[system][mysql][createTable] create table `administrators`")
	}
	if !db.HasTable(&model.ClassifyRecord{}) {
		db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(&model.ClassifyRecord{})
		db.Model(&model.ClassifyRecord{}).AddForeignKey("user_id", "users(id)", "RESTRICT", "RESTRICT")
		db.Model(&model.ClassifyRecord{}).AddForeignKey("image_id", "images(id)", "RESTRICT", "RESTRICT")
		db.Model(&model.ClassifyRecord{}).AddForeignKey("garbage_type_id", "garbage_types(id)", "RESTRICT", "RESTRICT")
		log.Print("[system][mysql][createTable] create table `classify_records`")
	}
	if !db.HasTable(&model.GarbageDetail{}) {
		db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(&model.GarbageDetail{})
		db.Model(&model.GarbageDetail{}).AddForeignKey("base_type", "garbage_types(id)", "RESTRICT", "RESTRICT")
		db.Model(&model.GarbageDetail{}).AddForeignKey("image_id", "images(id)", "RESTRICT", "RESTRICT")
		log.Print("[system][mysql][createTable] create table `garbage_details`")
	}
	if !db.HasTable(&model.GarbageType{}) {
		db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(&model.GarbageType{})
		db.Model(&model.GarbageType{}).AddForeignKey("parent_type_id", "garbage_types(id)", "RESTRICT", "RESTRICT")
		db.Model(&model.GarbageType{}).AddForeignKey("image_id", "images(id)", "RESTRICT", "RESTRICT")
		db.Model(&model.GarbageType{}).AddForeignKey("garbage_detail_id", "garbage_details(id)", "RESTRICT", "RESTRICT")
		log.Print("[system][mysql][createTable] create table `garbage_types`")
	}
	if !db.HasTable(&model.Image{}) {
		db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(&model.Image{})
		log.Print("[system][mysql][createTable] create table `images`")
	}
	if !db.HasTable(&model.Recycle{}) {
		db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(&model.Recycle{})
		db.Model(&model.Recycle{}).AddForeignKey("user_id", "users(id)", "RESTRICT", "RESTRICT")
		db.Model(&model.Recycle{}).AddForeignKey("classify_record_id", "classify_records(id)", "RESTRICT", "RESTRICT")
		log.Print("[system][mysql][createTable] create table `images`")
	}
	if !db.HasTable(&model.User{}) {
		db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(&model.User{})
		db.Model(&model.User{}).AddForeignKey("avatar_image_id", "images(id)", "RESTRICT", "RESTRICT")
		log.Print("[system][mysql][createTable] create table `users`")
	}
	if !db.HasTable(&model.SetAppend{}) {
		db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(&model.SetAppend{})
		db.Model(&model.SetAppend{}).AddForeignKey("image_id", "images(id)", "RESTRICT", "RESTRICT")
		db.Model(&model.SetAppend{}).AddForeignKey("record_id", "classify_records(id)", "RESTRICT", "RESTRICT")
		log.Print("[system][mysql][createTable] create table `users`")
	}

	log.Print("[system][mysql] create table success!")
}
