package model

import "time"

type AdImage struct {
	AdImageId    uint64     `json:"intAdImageId" form:"intAdImageId"`
	AdId         uint64     `json:"intAdId" form:"intAdId"`
	Directory    string     `json:"szDirectory" form:"szDirectory"`
	Filename     string     `json:"szFilename" form:"szFilename"`
	CreatedDate  time.Time  `json:"dtmCreatedDate" form:"dtmCreatedDate"`
	ModifiedDate *time.Time `json:"dtmModifiedDate" form:"dtmModifiedDate"`
}
