package model

import "time"

type ReportGroup struct {
	ReportGroupId   uint64     `json:"intReportGroupId" form:"intReportGroupId"`
	ReportGroupName string     `json:"szReportGroupName" form:"szReportGroupName"`
	Annotation      *string    `json:"szAnnotation" form:"szAnnotation"`
	CreatedDate     time.Time  `json:"dtmCreatedDate" form:"dtmCreatedDate"`
	ModifiedDate    *time.Time `json:"dtmModifiedDate" form:"dtmModifiedDate"`
}
