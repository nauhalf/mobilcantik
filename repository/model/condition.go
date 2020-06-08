package model

import "time"

type Condition struct {
	ConditionId   uint64     `json:"intConditionId" form:"intConditionId"`
	ConditionName string     `json:"szConditionName" form:"szConditionName"`
	Annotation    *string    `json:"szAnnotation" form:"szAnnotation"`
	CreatedDate   time.Time  `json:"dtmCreatedDate" form:"dtmCreatedDate"`
	ModifiedDate  *time.Time `json:"dtmModifiedDate" form:"dtmModifiedDate"`
}
