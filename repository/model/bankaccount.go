package model

import "time"

type BankAccount struct {
	BankAccountId     uint64     `json:"intBankAccountId" form:"intBankAccountId"`
	BankAccountNumber string     `json:"szProvinceId" form:"szProvinceId"`
	BankCode          string     `json:"szBankCode" form:"szBankCode"`
	BankAccountName   string     `json:"szBankAccountName" form:"szBankAccountName"`
	IsActive          bool       `json:"isActive" form:"isActive"`
	CreatedDate       time.Time  `json:"dtmCreatedDate" form:"dtmCreatedDate"`
	ModifiedDate      *time.Time `json:"dtmModifiedDate" form:"dtmModifiedDate"`
}
