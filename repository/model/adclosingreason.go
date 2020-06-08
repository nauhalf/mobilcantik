package model

import "time"

type AdClosingReason struct {
	AdClosingReasonId uint64     `json:"intAdClosingReasonId" form:"intAdClosingReasonId"`
	ReasonName        string     `json:"szReasonName" form:"szReasonName"`
	Annotation        *string    `json:"szAnnotation" form:"szAnnotation"`
	CreatedDate       time.Time  `json:"dtmCreatedDate" form:"dtmCreatedDate"`
	ModifiedDate      *time.Time `json:"dtmModifiedDate" form:"dtmModifiedDate"`
}
