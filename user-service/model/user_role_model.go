package model

import "time"

type UserRole struct {
	ID        uint      `gorm:"primaryKey:autoIncrement" json:"id"`
	UserID    uint      `gorm:"not null" json:"user_id"`
	RoleID    uint      `gorm:"not null" json:"role_id"`
	User      User      `gorm:"foreignKey:UserID"`
	Role      Role      `gorm:"foreignKey:RoleID"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Tabler interface {
	TableName() string
}

func (UserRole) TableName() string {
	return "user_role"
}