package placemodel

import (
	"errors"
	"go01-airbnb/pkg/common"
	"strings"
)

const EntityName = "place"

type Place struct {
	common.SQLModel
	OwnerId       int                `json:"-" gorm:"column:owner_id"`
	Owner         *common.SimpleUser `json:"owner" gorm:"preload:false"`
	CityId        int                `json:"cityId" gorm:"column:city_id"`
	Name          string             `json:"name" gorm:"column:name"`
	Address       string             `json:"address" gorm:"column:address"`
	Cover         *common.Images     `json:"cover" gorm:"column:cover"`
	Lat           float64            `json:"lat" gorm:"column:lat"`
	Lng           float64            `json:"lng" gorm:"column:lng"`
	PricePerNight float64            `json:"pricePerNight" gorm:"column:price_per_night"`
}

func (Place) TableName() string {
	return "places"
}

func (p *Place) Validate() error {
	p.Name = strings.TrimSpace(p.Name)
	if p.Name == "" {
		return ErrNameIsEmpty
	}

	p.Address = strings.TrimSpace(p.Address)
	if p.Address == "" {
		return ErrAddressIsEmpty
	}

	return nil
}

var (
	ErrNameIsEmpty = common.NewCustomError(
		errors.New("place name can't be blank"),
		"place name can't be blank",
	)

	ErrAddressIsEmpty = common.NewCustomError(
		errors.New("place address can't be blank"),
		"place address can't be blank",
	)
)
