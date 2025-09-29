package main

import "gorm.io/gorm"

// structure & ORM

type User struct{
	gorm.Model
	Email string `gorm:"unique`
	Password string
}