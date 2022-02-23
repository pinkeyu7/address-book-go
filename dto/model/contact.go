package model

import "time"

type Contact struct {
	Id        int       `xorm:"pk autoincr BIGINT(20)" json:"id"`
	Name      string    `xorm:"not null default '' comment('name') VARCHAR(25)" json:"name"`
	Email     string    `xorm:"not null default '' comment('email') VARCHAR(40)" json:"email"`
	Phone     string    `xorm:"not null default '' comment('phone') VARCHAR(25)" json:"phone"`
	IsDisable bool      `xorm:"not null is_disable" json:"is_disable"`
	CreatedAt time.Time `xorm:"not null created DATETIME" json:"created_at"`
	UpdatedAt time.Time `xorm:"not null updated DATETIME" json:"updated_at"`
}
