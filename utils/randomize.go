package utils

import (
	"math/rand"
	"time"
)

// GenRandName will generate a random set of names
// to be used in unit test purposes.
func GenRandName(n int) string {
	const alphabet = "abcdefghijklmnopqrstuvwxyzABDCEFGHIJKLMNOPQRSTUVWXYZ"
	
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	buf := make([]byte, n)
	
	for i := range buf {
		buf[i] = alphabet[seededRand.Intn(len(alphabet))]
	}
	
	return string(buf)
}

// GenRandMoney will generate a random amount of money determined
// by min and max to be used in unit test purposes.
func GenRandMoney(min, max float64) float64 {
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	r := min + seededRand.Float64() * (max - min)
	return r
}

// GenRandCurrency will return a random currency based on
// the input of its parameter to be used in unit test purposes.
func GenRandCurrency(currs ...string) string {
	seed := rand.New(rand.NewSource(time.Now().UnixNano()))
	
	return currs[seed.Intn(len(currs))]
}