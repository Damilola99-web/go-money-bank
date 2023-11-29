package util

import (
	"math/rand"
	"strings"
	"time"
)

const alphapet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())

}

// returns a random int64 between min and max
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// returns a random string of length n
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphapet)
	for i := 0; i < n; i++ {
		sb.WriteByte(alphapet[rand.Intn(k)])
	}
	return sb.String()
}

// returns a random owner name
func RandomOwner() string {
	return RandomString(6)
}

// returns a random amount of money
func RandomMoney() int64 {
	return RandomInt(100, 1000)

}

// returns a random currency code
func RandomCurrency() string {
	currencies := []string{"EUR", "USD", "CAD"}
	n := len(currencies)
	return currencies[rand.Intn(n)]
}
