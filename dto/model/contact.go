package model

import "time"

type Contact struct {
	Id        int       `xorm:"pk autoincr BIGINT(20)" json:"id"`
	Name      string    `xorm:"not null default '' comment('name') VARCHAR(64)" json:"name"`
	Email     string    `xorm:"not null default '' comment('email') VARCHAR(64)" json:"email"`
	Phone     string    `xorm:"not null default '' comment('phone') VARCHAR(20)" json:"phone"`
	Gender    *int      `xorm:"not null default 0 comment('gender, 0 = Male, 1 = Female, 2 = Secret') TINYINT(4)" json:"gender"`
	IsDisable bool      `xorm:"not null is_disable" json:"is_disable"`
	CreatedAt time.Time `xorm:"not null created DATETIME" json:"created_at"`
	UpdatedAt time.Time `xorm:"not null updated DATETIME" json:"updated_at"`
}
