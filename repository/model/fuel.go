package model

import "time"

type Fuel struct {
	FuelId       uint64     `json:"intFuelId" form:"intFuelId"`
	FuelName     string     `json:"szFuelName" form:"szFuelName"`
	Annotation   *string    `json:"szAnnotation" form:"szAnnotation"`
	CreatedDate  time.Time  `json:"dtmCreatedDate" form:"dtmCreatedDate"`
	ModifiedDate *time.Time `json:"dtmModifiedDate" form:"dtmModifiedDate"`
}
