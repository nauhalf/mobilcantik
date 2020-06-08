package model

import (
	"time"
)

type User struct {
	UserId       uint64     `json:"intUserId" form:"intUserId"`
	Username     string     `json:"szUsername" form:"szUsername"`
	Fullname     string     `json:"szFullname" form:"szFullname"`
	Email        string     `json:"szEmail" form:"szEmail"`
	Password     string     `json:"szPassword" form:"szPassword"`
	IsDefault    bool       `json:"isDefault" form:"isDefault"`
	CreatedDate  time.Time  `json:"dtmCreatedDate" form:"dtmCreatedDate"`
	ModifiedDate *time.Time `json:"dtmModifiedDate" form:"dtmModifiedDate"`
}
