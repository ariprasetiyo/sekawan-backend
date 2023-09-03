package unittest

import (
	"sekawan-backend/app/main/util"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncryptDescypt(t *testing.T) {
	planText := "ari_prasetiyo"
	chiperText := util.EncryptAES256(planText)
	resultDescypted := util.DecryptAES256(chiperText)
	assert.Equal(t, planText, resultDescypted)
	println("planText", planText)
	println("chiperText", chiperText)
}
