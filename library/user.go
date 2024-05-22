package library

type UserHeader struct {
	UserID     uint64     `header:"x-user-id" binding:"required"`
	Status     UserStatus `header:"x-user-status" binding:"required"`
	IsVerified uint8      `header:"is_verified" binding:"required"`
	Gender     Gender     `header:"x-user-gender" binding:"required"`
	CityID     uint64     `header:"x-user-city-id" binding:"required"`
}
