package randomutil

import "testing"

func Test_Utils_Random_CreateString(t *testing.T) {
	randomStringTable := map[string]bool{}
	const iteration = 100
	const lenRandomString = 10

	for i := 0; i < iteration; i += 1 {
		randomString, err := Random.CreateString(lenRandomString)

		if err != nil {
			t.Fatalf(err.Error())
			return
		}

		if randomStringTable[randomString] {
			t.Fatalf("Identical random values")
			return
		}

		randomStringTable[randomString] = true
	}
}
