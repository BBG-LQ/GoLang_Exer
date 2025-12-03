package utils

import "github.com/jinzhu/gorm"


type DBUtil struct {
}

func (u DBUtil) Connect() *gorm.DB {

	dsn := "root:123456@tcp(127.0.0.1:3306)/test001?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open("mysql", dsn)

	if err != nil {
		panic(err)
	}

	return db

}

func (u DBUtil) Close(db *gorm.DB) {
	db.Close()
}
