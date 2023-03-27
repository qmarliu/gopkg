package mysql

import (
	"fmt"
	"time"

	"github.com/qmarliu/gopkg/log"
	"github.com/qmarliu/gopkg/params"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Init(mysqlConfig params.Mysql) (*gorm.DB, error) {
	if mysqlConfig.DbName == "" {
		return nil, fmt.Errorf("the name of the database is not specified")
	}
	// dsn := mysqlConfig.Username + ":" + mysqlConfig.Password + "@tcp(" +
	// 	mysqlConfig.Path + ")/" + mysqlConfig.DbName + "?" + mysqlConfig.Config
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?%s",
		mysqlConfig.Username, mysqlConfig.Password, mysqlConfig.Path, "mysql", mysqlConfig.Config)

	var db *gorm.DB
	var err1 error
	db, err := gorm.Open(mysql.Open(dsn), nil)
	if err != nil {
		fmt.Println("Open failed ", err.Error(), dsn)
	}
	if err != nil {
		time.Sleep(time.Duration(30) * time.Second)
		db, err1 = gorm.Open(mysql.Open(dsn), nil)
		if err1 != nil {
			fmt.Println("Open failed ", err1.Error(), dsn)
			panic(err1.Error())
		}
	}

	//Check the database and table during initialization
	sql := fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s default charset utf8 COLLATE utf8_general_ci;", mysqlConfig.DbName)
	fmt.Println("exec sql: ", sql, " begin")
	err = db.Exec(sql).Error
	if err != nil {
		fmt.Println("Exec failed ", err.Error(), sql)
		return nil, err
	}
	dsn = fmt.Sprintf("%s:%s@tcp(%s)/%s?%s",
		mysqlConfig.Username, mysqlConfig.Password, mysqlConfig.Path, mysqlConfig.DbName, mysqlConfig.Config)
	cf := mysql.Config{
		DSN:                       dsn,   // DSN data source name
		DefaultStringSize:         512,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据版本自动配置
	}
	var level = logger.Info
	switch mysqlConfig.LogMode {
	case "silent":
		level = logger.Silent
	case "info":
		level = logger.Info
	case "warn":
		level = logger.Warn
	case "error":
		level = logger.Error
	default:
		level = logger.Info
	}

	gormCf := &gorm.Config{
		Logger: logger.New(
			log.GetLog(),
			logger.Config{
				SlowThreshold:             time.Second,
				LogLevel:                  level,
				IgnoreRecordNotFoundError: true,
				Colorful:                  true,
			},
		),
	}

	if db, err = gorm.Open(mysql.New(cf), gormCf); err != nil {
		return nil, err
	}
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	sqlDB.SetMaxIdleConns(mysqlConfig.MaxIdleConns)
	sqlDB.SetMaxOpenConns(mysqlConfig.MaxOpenConns)
	return db, err

}
