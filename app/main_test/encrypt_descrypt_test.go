package unittest

import (
	"sekawan-backend/app/main/util"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncryptDescypt(t *testing.T) {
	key := "jb(HH}=#jA=%6QK7"
	planText := "8ARSgYr1ofFRGJrxoAga"
	chiperText := util.EncryptAES256(key, planText)
	resultDescypted := util.DecryptAES256(key, chiperText)
	assert.Equal(t, planText, resultDescypted)
	println("planText", planText)
	println("chiperText", chiperText)
}
