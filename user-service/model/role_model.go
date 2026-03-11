package model

import "time"

type Role struct {
	ID        uint      `gorm:"primaryKey:autoIncrement" json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Users	 []User    `json:"users" gorm:"many2many:user_role;"`
}