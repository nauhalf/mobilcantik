package model

import "time"

type Testimonial struct {
	TestimonialId     uint64     `json:"intTestimonialId" form:"intTestimonialId"`
	AdId              uint64     `json:"intAdId" form:"intAdId"`
	Comment           string     `json:"szComment" form:"szComment"`
	AdClosingReasonId uint64     `json:"intAdClosingReasonId" form:"intAdClosingReasonId"`
	DtmDate           time.Time  `json:"dtmDate" form:"dtmDate"`
	IsActive          bool       `json:"isActive" form:"isActive"`
	CreatedDate       time.Time  `json:"dtmCreatedDate" form:"dtmCreatedDate"`
	ModifiedDate      *time.Time `json:"dtmModifiedDate" form:"dtmModifiedDate"`
}
