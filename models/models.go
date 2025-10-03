package models

import "time"

type User struct {
	ID            int       `json:"id" db:"id"`
	FirstName     *string   `json:"first_name" validate:"required,min=2,max=30" db:"first_name"`
	LastName      *string   `json:"last_name" validate:"required,min=2,max=30" db:"last_name"`
	Password      *string   `json:"password" validate:"required,min=6" db:"password"`
	Email         *string   `json:"email" validate:"required,email" db:"email"`
	Phone        *string   `json:"phone" validate:"required" db:"phone_no"`
	Token         *string   `json:"token" db:"token"`
	RefreshToken  *string   `json:"refresh_token" db:"refresh_token"`
	CreatedAt     time.Time `json:"created_at" db:"created_at"`
	UpdatedAt     time.Time `json:"updated_at" db:"updated_at"`
	AddressDetails []Address `json:"address"`
	CollegeDetails []CollegeSTD `json:"college"`
}

type CollegeSTD struct {
	CollegeID   int     `json:"id" db:"college_id"`
	UserID    int     `json:"user_id" db:"user_id"`   // foreign key to users.id
	CollegeName *string `json:"college_name" db:"college_name"`
	Country     *string `json:"country" db:"country"`
	State       *string `json:"state" db:"state"`
	City        *string `json:"city" db:"city"`
}

type  Address struct{
	AddressID int     `json:"id" db:"address_id"`
	UserID    int     `json:"user_id" db:"user_id"`   // foreign key to users.id
	Country    *string 		`json:"country" db:"country"`
	State  		   *string 		`json:"state" db:"state"`
	City   		   *string 		`json:"city" db:"city"`
}