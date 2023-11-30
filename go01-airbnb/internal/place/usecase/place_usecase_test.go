package placeusecase

import (
	"context"
	"errors"
	placemodel "go01-airbnb/internal/place/model"
	"go01-airbnb/pkg/common"
	"testing"
)

type mockPlaceRepository struct{}

// Giả lập hàm Create của Place Repository
func (mockPlaceRepository) Create(context context.Context, place *placemodel.Place) error {
	if place.Name == "New World" {
		return common.ErrDB(errors.New("something went wrong in DB"))
	}

	place.Id = 101
	return nil
}

func (mockPlaceRepository) Update(context context.Context, condition map[string]any, place *placemodel.Place) error {
	return nil
}

func (mockPlaceRepository) Delete(context context.Context, condition map[string]any) error {
	return nil
}

func (mockPlaceRepository) ListDataWithCondition(context context.Context, paging *common.Paging, filter *placemodel.Filter, keys ...string) ([]placemodel.Place, error) {
	return nil, nil
}

func (mockPlaceRepository) FindDataWithCondition(context context.Context, condition map[string]any, keys ...string) (*placemodel.Place, error) {
	return nil, nil
}

func TestCreatePlace(t *testing.T) {
	placeUC := NewPlaceUseCase(mockPlaceRepository{})

	// Test những trường hợp lỗi
	dataTable := []struct {
		Input    placemodel.Place
		Expected string
	}{
		{Input: placemodel.Place{Name: "", Address: "Nguyễn Huệ, Q1"}, Expected: placemodel.ErrNameIsEmpty.Error()},
		{Input: placemodel.Place{Name: "Candy Home", Address: ""}, Expected: placemodel.ErrAddressIsEmpty.Error()},
		{Input: placemodel.Place{Name: "New World", Address: "Lê Lai, Q1"}, Expected: "something went wrong in DB"},
	}

	for _, item := range dataTable {
		err := placeUC.CreatePlace(context.Background(), &item.Input)
		if err == nil || err.Error() != item.Expected {
			t.Errorf("create place - Input: %v, Expected: %v, Output: %v", item.Input, item.Expected, err)
		}
	}

	// Test trường hợp thành công, expect là không trả ra lỗi
	dataTest := placemodel.Place{Name: "Sweet Home", Address: "CMT8, Q3"}
	err := placeUC.CreatePlace(context.Background(), &dataTest)
	if err != nil {
		t.Errorf("create place - Input: %v, Output: %v", dataTest, err)
	}
}
