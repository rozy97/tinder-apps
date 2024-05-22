package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GetMD5Hash(t *testing.T) {

	t.Run("success case", func(t *testing.T) {
		hash := GetMD5Hash("12345" + "secret")
		assert.Equal(t, "d6bf7523a8407696bb9448d0d0fecca8", hash)
	})

}
