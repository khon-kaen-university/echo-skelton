package datamodels

import (
	"time"
)

// MainUsers structure for table: main_users
type MainUsers struct {
	Username  string    `gorm:"type:varchar(32);primary_key" json:"uusername"`
	FirstName  string    `gorm:"type:varchar(128);index" json:"first_name"`
	LastName  string    `gorm:"type:varchar(128);index" json:"last_name"`
	CreatedAt time.Time `gorm:"type:datetime;default:current_timestamp;index" json:"create_at"`
	UpdatedAt time.Time `gorm:"type:datetime;index" json:"updated_at"`
	DeletedAt time.Time `gorm:"type:datetime;index" json:"deleted_at"`
}
