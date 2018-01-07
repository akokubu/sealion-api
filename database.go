package main

import (
	"fmt"
	"net/url"
	"os"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB

func init() {
	var err error
	var datasource string
	if os.Getenv("DATABASE_URL") != "" {
		// for Heroku
		datasource = convert_datasource(os.Getenv("CLEARDB_DATABASE_URL"))
	} else {
		// for local
		datasource = "sealion:sealion@/sealion?charset=utf8&parseTime=true"
	}
	db, err = gorm.Open("mysql", datasource)
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&Task{})
}

// データソース文字列を変換
func convert_datasource(ds string) (result string) {
	url, _ := url.Parse(ds)
	result = fmt.Sprintf("%s@tcp(%s:3306)%s?parseTime=true", url.User.String(), url.Host, url.Path)
	return
}
