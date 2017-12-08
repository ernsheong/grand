package grand

import (
	"math/rand"
	"regexp"
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

func TestSetDefaultCharSet(t *testing.T) {
	randStr1 := GenerateRandomString(20)

	// There should not be any numbers in the string
	res := regexp.MustCompile("\\d").FindAllString(randStr1, -1)
	if len(res) != 0 {
		t.Fatal()
	}
}

func TestSetDifferentCharSet(t *testing.T) {
	gen := NewGenerator(CharSetBase62)

	randStr1 := gen.GenerateRandomString(20)

	// There should be at least one number in generated string
	res := regexp.MustCompile("\\d").FindAllString(randStr1, -1)
	if len(res) == 0 {
		t.Fatal()
	}
}
