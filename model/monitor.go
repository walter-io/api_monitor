package model

import (
	"apiMonitor/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

/**
 * 数据详情
 */
type Detail struct {
	RemoteAddr  string  `gorm:"column:ip"`
	RemoteUser  string  `gorm:"column:remote_user"`
	Time        string  `gorm:"column:request_time"`
	Method      string  `gorm:"column:method"`
	RequestUrl  string  `gorm:"column:request_url"`
	Protocol    string  `gorm:"column:protocol"`
	Status      int     `gorm:"column:status"`
	Size        float64 `gorm:"column:size"`
	OriginUrl   string  `gorm:"column:origin_url"`
	UserAgent   string  `gorm:"column:agent"`
	RequestTime float64 `gorm:"column:time"`
}

/**
 * 连接数据库
 */
func ConnectDb() gorm.DB {
	dsn := config.MysqlUsername + ":" + config.MysqlPassword + "@tcp(" + config.MysqlHost + ":" + config.MysqlPort + ")/" +
		config.MysqlDatabase + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return *db
}

/**
 * 插入数据
 */
func InsertDetail(row Detail) {
	db := ConnectDb()
	res := db.Create(row)
	if  res.Error != nil {
		panic(res.Error)
	}
}
