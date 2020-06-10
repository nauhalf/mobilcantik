package model

import "time"

type AdConfirmation struct {
	AdConfirmationId    uint64     `json:"intAdConfirmationId" form:"intAdConfirmationId"`
	AdId                uint64     `json:"intAdId" form:"intAdId"`
	AdBillId            uint64     `json:"intAdBillId" form:"intAdBillId"`
	BankAccountId       uint64     `json:"intBankAccountId" form:"intBankAccountId"`
	BankAccountNumber   string     `json:"szBankAccountNumber" form:"szBankAccountNumber"`
	BankName            string     `json:"szBankName" form:"szBankName"`
	BankBeneficiaryName string     `json:"szBankBeneficiaryName" form:"szBankBeneficiaryName"`
	Annotation          string     `json:"szAnnotation" form:"szAnnotation"`
	Date                time.Time  `json:"dtmDate" form:"dtmDate"`
	CreatedDate         time.Time  `json:"dtmCreatedDate" form:"dtmCreatedDate"`
	ModifiedDate        *time.Time `json:"dtmModifiedDate" form:"dtmModifiedDate"`
}

type AdConfirmationImage struct {
	AdConfirmationId uint64 `json:"intAdConfirmationId" form:"intAdConfirmationId"`
	Directory        string `json:"szDirectory" form:"szDirectory"`
	Filename         string `json:"szFilename" form:"szFilename"`
}
