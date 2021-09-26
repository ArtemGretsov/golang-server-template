package randomutil

import (
	"math/rand"
)

type random struct {
	Chars string
}

var Random = random{
	Chars: "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ",
}

func (r random) CreateString(length int) (string, error) {
	bytes := make([]byte, length)

	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}

	for i, b := range bytes {
		bytes[i] = r.Chars[b%byte(len(r.Chars))]
	}

	return string(bytes), nil
}