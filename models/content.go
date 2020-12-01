package models

import (
	"time"
)

type Content struct {
	AtticleId int       `xorm:"not null pk INT(11)"`
	Content   string    `xorm:"not null TEXT"`
	CreatedAt time.Time `xorm:"not null created default 'CURRENT_TIMESTAMP' TIMESTAMP"`
	UpdatedAt time.Time `xorm:"not null updated default 'CURRENT_TIMESTAMP' TIMESTAMP"`
}
