package model

import "time"

type AdVehicle struct {
	AdVehicleId    uint64     `json:"intAdVehicleId" form:"intAdVehicleId"`
	AdId           uint64     `json:"intAdId" form:"intAdId"`
	VehicleTypeId  uint64     `json:"intVehicleTypeId" form:"intVehicleTypeId"`
	VehicleBrandId uint64     `json:"intVehicleBrandId" form:"intVehicleBrandId"`
	Type           string     `json:"szType" form:"szType"`
	Model          string     `json:"szModel" form:"szModel"`
	ConditionId    uint64     `json:"intConditionId" form:"intConditionId"`
	ProductYear    int        `json:"intProductYear" form:"intProductYear"`
	FuelId         uint64     `json:"intFuelId" form:"intFuelId"`
	Color          string     `json:"szColor" form:"szColor"`
	Kilometer      int        `json:"intKilometer" form:"intKilometer"`
	Price          float64    `json:"decPrice" form:"decPrice"`
	Annotation     *string    `json:"szAnnotation" form:"szAnnotation"`
	CreatedDate    time.Time  `json:"dtmCreatedDate" form:"dtmCreatedDate"`
	ModifiedDate   *time.Time `json:"dtmModifiedDate" form:"dtmModifiedDate"`
}
