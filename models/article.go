package models

import (
	"time"
)

type Article struct {
	Id        int       `xorm:"not null pk autoincr INT(11)"`
	Type      int       `xorm:"not null default 0 TINYINT(4)"`
	Title     string    `xorm:"not null default '' VARCHAR(200)"`
	SubTitle  string    `xorm:"not null default '' VARCHAR(500)"`
	IsDel     int       `xorm:"not null default 0 TINYINT(4)"`
	Url       string    `xorm:"not null default '' VARCHAR(200)"`
	CreatedAt time.Time `xorm:"not null created default 'CURRENT_TIMESTAMP' TIMESTAMP"`
	UpdatedAt time.Time `xorm:"not null updated default 'CURRENT_TIMESTAMP' TIMESTAMP"`
}
