package model

import "time"

type Ad struct {
	AdId            uint64     `json:"intAdId" form:"intAdId"`
	Email           string     `json:"szEmail" form:"szEmail"`
	Title           string     `json:"szTitle" form:"szTitle"`
	Sender          string     `json:"szSender" form:"szSender"`
	CityId          string     `json:"szCityId" form:"szCityId"`
	PhoneNumber     string     `json:"szPhoneNumber" form:"szPhoneNumber"`
	Annotation      string     `json:"szAnnotation" form:"szAnnotation"`
	IsActive        bool       `json:"isActive" form:"isActive"`
	AdTypeId        uint64     `json:"intAdTypeId" form:"intAdTypeId"`
	DtmPosted       *time.Time `json:"dtmPosted" form:"dtmPosted"`
	DtmConfirmation *time.Time `json:"dtmConfirmation" form:"dtmConfirmation"`
	DtmExpired      *time.Time `json:"dtmExpired" form:"dtmExpired"`
	AdStatusTypeId  uint64     `json:"intAdStatusTypeId" form:"intAdStatusTypeId"`
	Hits            uint32     `json:"intHits" form:"intHits"`
	IsExpired       bool       `json:"isExpired" form:"isExpired"`
	CreatedDate     time.Time  `json:"dtmCreatedDate" form:"dtmCreatedDate"`
	ModifiedDate    *time.Time `json:"dtmModifiedDate" form:"dtmModifiedDate"`
}
