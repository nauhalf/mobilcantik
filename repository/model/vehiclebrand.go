package model

import "time"

type VehicleBrand struct {
	VehicleBrandId   uint64     `json:"intVehicleBrandId" form:"intVehicleBrandId"`
	VehicleBrandName string     `json:"szBrandName" form:"szBrandName"`
	VehicleTypeId    uint64     `json:"intVehicleTypeId" form:"intVehicleTypeId"`
	Count            uint64     `json:"intCount" form:"intCount"`
	Annotation       *string    `json:"szAnnotation" form:"szAnnotation"`
	CreatedDate      time.Time  `json:"dtmCreatedDate" form:"dtmCreatedDate"`
	ModifiedDate     *time.Time `json:"dtmModifiedDate" form:"dtmModifiedDate"`
}
