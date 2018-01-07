package main

import "github.com/jinzhu/gorm"

type Task struct {
	gorm.Model
	Title   string
	Project string
	Done    bool
}
