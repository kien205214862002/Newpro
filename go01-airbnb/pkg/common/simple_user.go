package common

// Struct thể hiện những thông tin có thể public của 1 user
type SimpleUser struct {
	SQLModel
	FirstName string `json:"firstName" gorm:"column:first_name"`
	LastName  string `json:"lastName" gorm:"column:last_name"`
}

func (SimpleUser) TableName() string {
	return "users"
}
