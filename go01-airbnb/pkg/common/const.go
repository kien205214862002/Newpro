package common

const (
	DBTypeUser  = 1
	DBTypePlace = 2
)

type Requester interface {
	GetUserId() int
	GetUserEmail() string
	GetUserRole() string
}
