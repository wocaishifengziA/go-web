package database

import (
	"database/sql"
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const DbName = "ccc"

var DB *gorm.DB

func DbConnect() error {
	dsn := "root:123456@tcp(192.168.2.125:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local"
	dsnDefault := fmt.Sprintf(dsn, "mysql")
	db, err := gorm.Open(mysql.Open(dsnDefault), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("open dsn %s err: %s", "mysql", err.Error())
	}
	var count int
	err = db.Raw("SELECT count(1) FROM information_schema.SCHEMATA where SCHEMA_NAME=@name", sql.Named("name", DbName)).Scan(&count).Error
	if err != nil {
		return fmt.Errorf("db raw err: %s", err.Error())
	}

	if count == 0 {
		sql := "CREATE DATABASE " + DbName
		err = db.Exec(sql).Error
		if err != nil {
			return fmt.Errorf("create db %s err: %s", DbName, err.Error())
		}
	}

	dsnDb := fmt.Sprintf(dsn, DbName)
	DB, err = gorm.Open(mysql.Open(dsnDb), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("open dsn %s err: %s", dsnDb, err.Error())
	}
	return nil
}

func DbPing() {
	for {
		sqlDB, _ := DB.DB()
		err := sqlDB.Ping()
		if err != nil {
			fmt.Println("DbPing err: ", err.Error())
			err = DbConnect()
			if err != nil {
				fmt.Println("Db Connect err: ", err.Error())
			}
		}
		time.Sleep(time.Second * 10)
	}
}

func CreateTableIfNotExists(modelType interface{}) error {
	if !DB.Migrator().HasTable(modelType) {
		return DB.Migrator().CreateTable(modelType)
	}
	return nil
}
