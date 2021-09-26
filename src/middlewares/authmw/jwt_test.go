package authmw

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_AuthGuard_Create_And_Parse_JWT(t *testing.T) {
	inJWTPayload := JWTPayload{
		ID: 1,
		Name: "name",
		Login: "login",
	}

	token, err := CreateJWT(inJWTPayload)
	assert.Nil(t, err)

	outJWTPayload, err := ParseJWT(token)
	assert.Nil(t, err)

	assert.Equal(t, inJWTPayload, outJWTPayload, "payload is different")
}
