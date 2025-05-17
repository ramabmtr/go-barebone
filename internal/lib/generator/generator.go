package generator

import (
	"bytes"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

const (
	Uppercase    = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	Lowercase    = "abcdefghijklmnopqrstuvwxyz"
	Alphabetic   = Uppercase + Lowercase
	Numeric      = "0123456789"
	Alphanumeric = Alphabetic + Numeric
	Symbols      = "`" + `~!@#$%^&*()-_+={}[]|\;:"<>,./?`
	CompleteChar = Alphanumeric + Symbols
	Hex          = Numeric + "abcdef"
)

var seededRand = rand.New(rand.NewSource(time.Now().UnixNano()))

func RandomNumber(length int) int {
	s := RandomString(length, Numeric)
	i, _ := strconv.Atoi(s)
	return i
}

func RandomString(length int, charset string, prefixes ...string) string {
	var s bytes.Buffer
	for _, prefix := range prefixes {
		s.WriteString(prefix)
	}

	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}

	s.Write(b)

	return s.String()
}

func GenerateCacheKey(keys ...string) string {
	divider := "::"
	key := "cache" + divider
	for _, s := range keys {
		key = key + s + divider
	}
	return strings.TrimSuffix(key, divider)
}
