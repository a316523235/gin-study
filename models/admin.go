package models

import (
	"time"
)

type Admin struct {
	Id        int       `xorm:"not null pk autoincr INT(11)"`
	Username  string    `xorm:"not null default '' VARCHAR(20)"`
	Pwd       string    `xorm:"not null default '' VARCHAR(50)"`
	CreatedAt time.Time `xorm:"not null created default 'CURRENT_TIMESTAMP' TIMESTAMP"`
	UpdatedAt time.Time `xorm:"not null updated default 'CURRENT_TIMESTAMP' TIMESTAMP"`
}
