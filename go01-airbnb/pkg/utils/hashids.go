package utils

import "github.com/speps/go-hashids/v2"

type Hasher struct {
	HashID *hashids.HashID
}

func NewHashIds(salt string, minLength int) *Hasher {
	hd := hashids.NewData()
	hd.Salt = salt
	hd.MinLength = minLength

	h, _ := hashids.NewWithData(hd)

	return &Hasher{h}
}

func (h *Hasher) Encode(id, dbType int) string {
	str, _ := h.HashID.Encode([]int{id, dbType})
	return str
}

func (h *Hasher) Decode(str string) (int, error) {
	ids, err := h.HashID.DecodeWithError(str)
	if err != nil {
		return 0, err
	}

	return ids[0], nil
}
