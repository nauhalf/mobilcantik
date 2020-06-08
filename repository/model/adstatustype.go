package model

import "time"

type AdStatusType struct {
	AdStatusTypeId   uint64     `json:"intAdStatusTypeId" form:"intAdStatusTypeId"`
	AdStatusTypeName string     `json:"szAdStatusTypeName" form:"szAdStatusTypeName"`
	CreatedDate      time.Time  `json:"dtmCreatedDate" form:"dtmCreatedDate"`
	ModifiedDate     *time.Time `json:"dtmModifiedDate" form:"dtmModifiedDate"`
}
