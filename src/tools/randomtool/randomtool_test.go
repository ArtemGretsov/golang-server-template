package randomtool

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_RandomTool_CreateString(t *testing.T) {
	randomStringTable := map[string]bool{}
	const iteration = 100
	const lenRandomString = 10

	for i := 0; i < iteration; i += 1 {
		randomString, err := CreateString(lenRandomString)

		assert.Nil(t, err)
		assert.False(t, randomStringTable[randomString],"Identical random values")

		randomStringTable[randomString] = true
	}
}
