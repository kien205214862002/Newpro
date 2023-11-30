package placemodel

type Filter struct {
	OwnerId int `json:"ownerId" form:"ownerId"`
	CityId  int `json:"cityId" form:"cityId"`
}
