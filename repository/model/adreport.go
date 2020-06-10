package model

import "time"

type AdReport struct {
	AdReportId    uint64     `json:"intAdReportId" form:"intAdReportId"`
	AdId          uint64     `json:"intAdId" form:"intAdId"`
	ReportTypeId  uint64     `json:"intReportTypeId " form:"intReportTypeId "`
	Annotation    string     `json:"szAnnotation" form:"szAnnotation"`
	ReportTime    time.Time  `json:"dtmReportTime" form:"dtmReportTime"`
	ReporterName  string     `json:"szReporterName" form:"szReporterName"`
	ReporterEmail string     `json:"szReporterEmail" form:"szReporterEmail"`
	CreatedDate   time.Time  `json:"dtmCreatedDate" form:"dtmCreatedDate"`
	ModifiedDate  *time.Time `json:"dtmModifiedDate" form:"dtmModifiedDate"`
}
