package placemodel

import "testing"

type testData struct {
	Input    Place
	Expected error
}

func TestPlaceValidate(t *testing.T) {
	data := []testData{
		{Input: Place{Name: "", Address: "Nguyễn Huệ, Q1"}, Expected: ErrNameIsEmpty},
		{Input: Place{Name: "New World", Address: ""}, Expected: ErrAddressIsEmpty},
		{Input: Place{Name: "New World", Address: "Lê Lai, Q1"}, Expected: nil},
	}

	for _, item := range data {
		err := item.Input.Validate()

		if err != item.Expected {
			t.Errorf("Validate place: Input %v, Expect: %v, Output: %v", item.Input, item.Expected, err)
		}
	}
}
