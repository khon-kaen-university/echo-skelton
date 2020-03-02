package datamodels

import (
	"time"
)

// MainUsers structure for table: main_users
type MainUsers struct {
	Username  string    `gorm:"type:varchar(128);primary_key" json:"uusername"`
	FirstName  string    `gorm:"type:varchar(128);primary_key" json:"first_name"`
	LastName  string    `gorm:"type:varchar(128);primary_key" json:"last_name"`
	CreatedAt time.Time `gorm:"type:datetime;default:current_timestamp;index" json:"create_at"`
	UpdatedAt time.Time `gorm:"type:datetime;index" json:"updated_at"`
	DeletedAt time.Time `gorm:"type:datetime;index" json:"deleted_at"`
}
