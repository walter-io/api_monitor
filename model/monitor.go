package model

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

/**
 * 数据详情
 */
type Detail struct { // todo 数据库字段要优化
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
	// TODO 连接数据库，哪里设置表？
	dsn := "root:123456@tcp(192.168.0.235:3306)/lara?charset=utf8mb4&parseTime=True&loc=Local"
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
	fmt.Printf("Error: %v, RowsAffected: %v\n", res.Error, res.RowsAffected)
}
