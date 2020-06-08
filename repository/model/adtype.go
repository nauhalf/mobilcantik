package model

import "time"

type AdType struct {
	AdTypeId     uint64     `json:"intAdTypeId" form:"intAdTypeId"`
	AdTypeName   string     `json:"szAdTypeName" form:"szAdTypeName"`
	Annotation   *string    `json:"szAnnotation" form:"szAnnotation"`
	CreatedDate  time.Time  `json:"dtmCreatedDate" form:"dtmCreatedDate"`
	ModifiedDate *time.Time `json:"dtmModifiedDate" form:"dtmModifiedDate"`
}
