package model

type Province struct {
	ProvinceId   string  `json:"szProvinceId" form:"szProvinceId"`
	ProvinceName string  `json:"szProvinceName" form:"szProvinceName"`
	Annotation   *string `json:"szAnnotation" form:"szAnnotation"`
}
