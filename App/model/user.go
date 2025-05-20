package model

import "time"

type User struct {
	Id          int64     `gorm:"column:id;primary_key;AUTO_INCREMENT;NOT NULL"`
	Password    string    `gorm:"column:password"`
	Name        string    `gorm:"column:name"`
	CreatedTime time.Time `gorm:"column:created_time"`
	UpdateTime  time.Time `gorm:"column:update_time"`
	Telephone   string    `gorm:"column:telephone;default:NULL"`
	Token       string    `gorm:"column:token;default:NULL"`
}

func (u *User) TableName() string {
	return "user"
}
