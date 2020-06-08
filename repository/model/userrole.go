package model

import (
	"time"
)

type UserRole struct {
	UserId       uint64     `json:"intUserId" form:"intUserId"`
	RoleId       string     `json:"szRoleId" form:"szRoleId"`
	CreatedDate  time.Time  `json:"dtmCreatedDate" form:"dtmCreatedDate"`
	ModifiedDate *time.Time `json:"dtmModifiedDate" form:"dtmModifiedDate"`
}
