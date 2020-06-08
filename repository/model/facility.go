package model

import "time"

type Facility struct {
	FacilityId    uint64     `json:"intFacilityId" form:"intFacilityId"`
	FacilityName  string     `json:"szFacilityName" form:"szFacilityName"`
	VehicleTypeId uint64     `json:"intVehicleTypeId" form:"intVehicleTypeId"`
	Annotation    *string    `json:"szAnnotation" form:"szAnnotation"`
	CreatedDate   time.Time  `json:"dtmCreatedDate" form:"dtmCreatedDate"`
	ModifiedDate  *time.Time `json:"dtmModifiedDate" form:"dtmModifiedDate"`
}
