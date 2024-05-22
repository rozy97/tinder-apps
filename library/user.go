package library

type UserHeader struct {
	UserID     uint64     `header:"x-user-id" binding:"required"`
	Status     UserStatus `header:"x-user-status" binding:"required"`
	IsVerified uint8      `json:"is_verified" binding:"required"`
}
