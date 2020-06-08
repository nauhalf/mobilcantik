package model

import "time"

type Partner struct {
	PartnerId    uint64     `json:"intPartnerId" form:"intPartnerId"`
	Title        string     `json:"szTitle" form:"szTitle"`
	Url          string     `json:"szUrl" form:"szUrl"`
	Annotation   *string    `json:"szAnnotation" form:"szAnnotation"`
	CreatedDate  time.Time  `json:"dtmCreatedDate" form:"dtmCreatedDate"`
	ModifiedDate *time.Time `json:"dtmModifiedDate" form:"dtmModifiedDate"`
}
