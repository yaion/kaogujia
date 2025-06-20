package middleware

import (
	"fmt"
	"kaogujia/pkg/config"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 全局数据库连接实例
var DB *gorm.DB

func InitMySQL() error {
	database := config.Get().Database
	fmt.Println("database:", database)
	dns := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",
		database.User, database.Password, database.Host, database.Port, database.DbName, database.Charset)
	fmt.Println(dns)
	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{})
	if err != nil {
		return err
	}

	// 配置连接池
	sqlDB, err := db.DB()
	if err != nil {
		return err
	}
	sqlDB.SetMaxIdleConns(database.MaxIdleConns)
	sqlDB.SetMaxOpenConns(database.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(time.Duration(int(database.ConnMaxLifetime)))

	DB = db
	return nil
}

// 获取数据库连接
func GetDB() *gorm.DB {
	return DB
}
