package model

import "time"

type ReportType struct {
	ReportTypeId   uint64     `json:"intReportTypeId" form:"intReportTypeId"`
	ReportGroupId  uint64     `json:"intReportGroupId" form:"intReportGroupId"`
	ReportTypeName string     `json:"szReportTypeName" form:"szReportTypeName"`
	Annotation     *string    `json:"szAnnotation" form:"szAnnotation"`
	CreatedDate    time.Time  `json:"dtmCreatedDate" form:"dtmCreatedDate"`
	ModifiedDate   *time.Time `json:"dtmModifiedDate" form:"dtmModifiedDate"`
}
