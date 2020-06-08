package model

import "time"

type VehicleType struct {
	VehicleTypeId   uint64     `json:"intVehicleTypeId" form:"intVehicleTypeId"`
	VehicleTypeName string     `json:"szVehicleTypeName" form:"szVehicleTypeName"`
	Annotation      *string    `json:"szAnnotation" form:"szAnnotation"`
	CreatedDate     time.Time  `json:"dtmCreatedDate" form:"dtmCreatedDate"`
	ModifiedDate    *time.Time `json:"dtmModifiedDate" form:"dtmModifiedDate"`
}
