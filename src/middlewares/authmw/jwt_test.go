package authmw

import (
	"testing"
)

func Test_AuthGuard_Create_And_Parse_JWT(t *testing.T) {
	inJWTPayload := JWTPayload{
		ID: 1,
		Name: "name",
		Login: "login",
	}

	token, err := CreateJWT(inJWTPayload)

	if err != nil {
		t.Fatal(err.Error())
		return
	}

	outJWTPayload, err := ParseJWT(token)

	if err != nil {
		t.Fatal(err.Error())
		return
	}

	if inJWTPayload != outJWTPayload {
		t.Fatal("payload is different")
	}
}
