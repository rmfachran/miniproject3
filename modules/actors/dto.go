package actors

import (
	"github.com/rmfachran/miniproject2/dto"
	"time"
)

type ActorParam struct {
	ID         uint   `gorm:"column:id"`
	Username   string `gorm:"column:username"`
	Password   string `gorm:"column:password"`
	RoleId     uint   `gorm:"column:role_id"`
	IsVerified string `gorm:"column:is_verified"`
	IsActive   string `gorm:"column:is_active"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type FindAdmin struct {
	dto.ResponseMeta
	Data ActorParam `json:"data"`
}

type SuccessCreate struct {
	dto.ResponseMeta
	Data ActorParam `json:"data"`
}

type SuccessLoginAdmin struct {
	dto.ResponseMeta
	Username string `json:"username"`
	Token    string `json:"token"`
}
type SuccessApproveAdmin struct {
	dto.ResponseMeta
	Data interface{} `json:"data"`
}

type SuccessGetCustomers struct {
	dto.ResponseMeta
	Data interface{} `json:"data"`
}

type SuccessUpdate struct {
	dto.ResponseMeta
	Data ActorParam `json:"data"`
}
