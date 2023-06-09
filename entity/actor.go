package entity

import "time"

type Actor struct {
	ID         uint   `gorm:"column:id"`
	Username   string `gorm:"column:username"`
	Password   string `gorm:"column:password"`
	RoleId     uint   `gorm:"column:role_id"`
	IsVerified string `gorm:"column:is_verified"`
	IsActive   string `gorm:"column:is_active"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
