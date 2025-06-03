package utils

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// TestGenRandName will test the random name generator
func TestGenRandName(t *testing.T) {
	ln := 10
	r := GenRandName(ln)
	
	require.NotEmpty(t, r)
	require.Len(t, r, ln)
	require.IsType(t, "", r)
}

// TestGenRandMoney will test the random money generator
func TestGenRandMoney(t *testing.T) {
	min := -10.0
	max := 10.0
	
	r := GenRandMoney(min, max)
	
	require.NotEmpty(t, r)
	require.IsType(t, float64(0), r)
	require.GreaterOrEqual(t, r, min)
	require.LessOrEqual(t, r, max)
}

// TestGenRandName will test the random currency generator
func TestGenRandCurrency(t *testing.T) {	
	r := GenRandCurrency("USD", "EUR", "DZD")
	
	require.NotEmpty(t, r)
	switch r {
		case "USD":
		return
		case "EUR":
		return
		case "DZD":
		return
		default:
		t.Errorf("FAILED: GenRandCurrecy failed to generate random currency from the set. wanted: %s, received: %s", "USD | EUR | DZD", r)
	}
}
