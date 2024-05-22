package config

import (
	"crypto/md5"
	"encoding/hex"
)

type (
	Gender int16
	Status int8
	Swipe  int8
)

const (
	GenderUnverified Gender = 0
	GenderMale       Gender = 1
	GenderFemale     Gender = 2

	StatusInactive Status = 0
	StatusActive   Status = 1

	SwipeDislike Swipe = 0
	SwipeLike    Swipe = 1
)

func GetMD5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}
