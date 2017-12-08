package grand

import (
	"math/rand"
	"testing"
	"time"
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func TestGeneratesDifferentStrings(t *testing.T) {
	randStr1 := GenerateRandomString(20)

	if len(randStr1) != 20 {
		t.Fatal()
	}

	randStr2 := GenerateRandomString(20)
	if randStr1 == randStr2 {
		t.Fatal()
	}

	randStr3 := GenerateRandomString(20)
	if randStr2 == randStr3 {
		t.Fatal()
	}
}
