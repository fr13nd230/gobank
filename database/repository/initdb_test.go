package repository

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

// TestNewDb will test db initialization function and
// assert its behaviour.
func TestNewDb(t *testing.T) {
	q, err := NewDb(os.Getenv("DB_PATH"))
	require.NoError(t, err)
	require.NotEmpty(t, q)
}
