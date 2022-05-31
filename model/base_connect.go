package model

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db_write *gorm.DB
var err error

var Db_read *gorm.DB

//func Init() {
//	write_ip := conf.Get("write_sql.ip")
//	write_password := conf.Get("write_sql.password")
//	write_port := conf.Get("write_sql.port")
//
//    dsn := "root:" + write_password + "@(" + write_ip + ":" + write_port + ")" + "/test?charset=utf8mb4&parseTime=True&loc=Local"
//	//dsn := "root:@(mysql-5.7-test:3306)/test_db?charset=utf8mb4&parseTime=True&loc=Local"
//	fmt.Println(dsn)
//	Db_write, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
//	if err != nil {
//		fmt.Println("write_sql err",err)
//	}
//
//
//	read_ip := conf.Get("read_sql.ip")
//	read_password := conf.Get("read_sql.password")
//	read_port := conf.Get("read_sql.port")
//
//	d := "root:" + read_password + "@(" + read_ip + ":" + read_port + ")" + "/test?charset=utf8mb4&parseTime=True&loc=Local"
//	//dsn := "root123456:@(mysql-5.7-test:3306)/test_db?charset=utf8mb4&parseTime=True&loc=Local"
//	fmt.Println(d)
//	Db_write, err = gorm.Open(mysql.Open(d), &gorm.Config{})
//	if err != nil {
//		fmt.Println("read_sql err",err)
//	}
//}

func Init() {
	dsn := "root:123456@(43.138.106.15:3306)/demo?charset=utf8mb4&parseTime=True&loc=Local"
	Db_write, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("write failed", err)
	}

	d := "root:123456@(116.62.150.88:3306)/demo?charset=utf8mb4&parseTime=True&loc=Local"
	Db_read, err = gorm.Open(mysql.Open(d), &gorm.Config{})

	if err != nil {
		fmt.Println("read failed", err)
	}

}
