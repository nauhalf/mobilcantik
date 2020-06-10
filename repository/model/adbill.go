package model

import "time"

type AdBill struct {
	AdBillId     uint64     `json:"intAdBillId" form:"intAdBillId"`
	AdId         uint64     `json:"intAdId" form:"intAdId"`
	Amount       float64    `json:"decAmount" form:"decAmount"`
	Balance      float64    `json:"decBalance" form:"decBalance"`
	CreatedDate  time.Time  `json:"dtmCreatedDate" form:"dtmCreatedDate"`
	ModifiedDate *time.Time `json:"dtmModifiedDate" form:"dtmModifiedDate"`
}
