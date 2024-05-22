package library

type (
	Gender     int16
	UserStatus int8
	Swipe      int8
)

const (
	GenderUnverified Gender = 0
	GenderMale       Gender = 1
	GenderFemale     Gender = 2

	StatusInactive UserStatus = 0
	StatusActive   UserStatus = 1

	SwipeDislike Swipe = 0
	SwipeLike    Swipe = 1
)
