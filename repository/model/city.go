package model

type City struct {
	CityId     string  `json:"szCityId" form:"szCityId"`
	ProvinceId string  `json:"szProvinceId" form:"szProvinceId"`
	CityName   string  `json:"szCityName" form:"szCityName"`
	Annotation *string `json:"szAnnotation" form:"szAnnotation"`
}
