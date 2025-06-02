package repository

import (
	"context"
	"testing"

	"github.com/fr13nd230/gobank/utils"
	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/require"
)

// CreateTransfer will create for us
// new transfer records, to use in later funcs.
func CreateTrasnfer(t *testing.T) Transfer {
	acc1 := CreateAccount(t)
	acc2 := CreateAccount(t)
	
	arg := NewTransferParams{
		FromAcc: acc1.ID,
		ToAcc: acc2.ID,
		Amount: utils.GenRandMoney(1, 100),
		Status: TransferstatusCreated,
	}
	
	r, err := queries.NewTransfer(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, r)
	
	require.NotZero(t, r.ID)
	require.NotZero(t, r.FromAcc)
	require.NotZero(t, r.ToAcc)
	require.NotZero(t, r.Amount)
	require.NotZero(t, r.Status)
	require.NotZero(t, r.CreatedAt)
	require.NotZero(t, r.UpdatedAt)
		
	require.Equal(t, r.FromAcc, acc1.ID)
	require.Equal(t, r.ToAcc, acc2.ID)
	require.Equal(t, r.Status, TransferstatusCreated)
	
	return r
}

// TestNewTransfer will test if the function
// will create new transfer without an error.
func TestNewTransfer(t *testing.T) {
	CreateTrasnfer(t)
}

// TestListTransfers will test if the function 
// will list the paginated transfers.
func TestListTransfers(t *testing.T) {
	limit := int32(10)
	offset := int32(1)
	arg := ListTransfersParams{
		Limit: limit,
		Offset: offset,
	}
	
	// Start transfers creation
	for i:=0; i<15; i++{
		CreateTrasnfer(t)
	}
	
	trs, err := queries.ListTransfers(context.Background(), arg)
	require.NoError(t, err)
	require.LessOrEqual(t, int(limit), len(trs))
	
	for _, tr := range trs{
		require.NotEmpty(t, tr)
		require.NotZero(t, tr.ID)
		require.NotZero(t, tr.FromAcc)
		require.NotZero(t, tr.ToAcc)
		require.NotZero(t, tr.Amount)
		require.NotZero(t, tr.Status)
		require.NotZero(t, tr.CreatedAt)
		require.NotZero(t, tr.UpdatedAt)
	}
}

// TestFindTransferById will try to find a transfer by its
// ID and then return it as a single object.
func TestFindTransferById(t *testing.T) {
	tr := CreateTrasnfer(t)

	find, err := queries.FindTransferById(context.Background(), tr.ID)
	
	require.NoError(t, err)
	require.NotEmpty(t, find)	
	require.NotZero(t, find.ID)
	require.NotZero(t, find.FromAcc)
	require.NotZero(t, find.ToAcc)
	require.NotZero(t, find.Amount)
	require.NotZero(t, find.Status)
	require.NotZero(t, find.CreatedAt)
	require.NotZero(t, find.UpdatedAt)
	require.Equal(t, tr, find)
}

// TestFindTransferByAcc will test if the function retrieves indeed
// the transfer only by account ID. 
func TestFindTransferByAcc(t *testing.T) {
	from, to := CreateAccount(t), CreateAccount(t)
	limit := int32(10)
	offset := int32(1)
	
	// Create transfers randomizing from and to
	i := 0
	for i < 20 {
		if i<10 {
			arg := NewTransferParams{
				FromAcc: from.ID,
				ToAcc: to.ID,
				Amount: utils.GenRandMoney(1, 100),
				Status: TransferstatusCreated,
			}
			_, err := queries.NewTransfer(context.Background(), arg)
			if err != nil {
				t.Errorf("FAILED: Test suite failed during creating a transfer: %v", err)
			}
			i++
		}
		arg := NewTransferParams{
			FromAcc: to.ID,
			ToAcc: from.ID,
			Amount: utils.GenRandMoney(1, 100),
			Status: TransferstatusCreated,
		}
		_, err := queries.NewTransfer(context.Background(), arg)
		if err != nil {
			t.Errorf("FAILED: Test suite failed during creating a transfer: %v", err)
		}
		i++
	}
	
	arg1 := FindTransfersByAccParams{
		FromAcc: from.ID,
		Limit: limit,
		Offset: offset,
	}
	arg2 := FindTransfersByAccParams{
		FromAcc: to.ID,
		Limit: limit,
		Offset: offset,
	}
	
	r1, err1 := queries.FindTransfersByAcc(context.Background(), arg1)
	r2, err2 := queries.FindTransfersByAcc(context.Background(), arg2)
	
	require.NoError(t, err1)
	require.NoError(t, err2)

	require.NotEmpty(t, r1)
	require.NotEmpty(t, r2)
	
	require.LessOrEqual(t, int(limit), len(r1))
	require.LessOrEqual(t, int(limit), len(r2))
	
	
	for _, tr := range r1 {
		require.NotEmpty(t, tr)
		require.NotZero(t, tr.ID)
		require.NotZero(t, tr.FromAcc)
		require.NotZero(t, tr.ToAcc)
		require.NotZero(t, tr.Amount)
		require.NotZero(t, tr.Status)
		require.NotZero(t, tr.CreatedAt)
		require.NotZero(t, tr.UpdatedAt)
	}
	
	for _, tr := range r2 {
		require.NotEmpty(t, tr)
		require.NotZero(t, tr.ID)
		require.NotZero(t, tr.FromAcc)
		require.NotZero(t, tr.ToAcc)
		require.NotZero(t, tr.Amount)
		require.NotZero(t, tr.Status)
		require.NotZero(t, tr.CreatedAt)
		require.NotZero(t, tr.UpdatedAt)
	}
}

// TestUpdateTransferById will create a new account and then attempts
// to update its status which originaly is of type TransferstatusCreated.
func TestUpdateTransferById(t *testing.T) {
	tr := CreateTrasnfer(t)
	
	arg := UpdateTransferByIdParams{
		ID: tr.ID,
		Status: TransferstatusProcessed,
	}
	
	r, err := queries.UpdateTransferById(context.Background(), arg)
	
	require.NoError(t, err)
	require.NotEmpty(t, r)
	
	require.NotEqual(t, r, tr.Status)
}

// TestDeleteTransferById will create a whole new transfer then
// uses its ID and test delete transfer func.
func TestDeleteTransferById(t *testing.T) {
	tr := CreateTrasnfer(t)
	
	err := queries.DeleteTransferById(context.Background(), tr.ID)
	require.NoError(t, err)
	
	// Try to find the error when deleted
	r, err := queries.FindTransferById(context.Background(), tr.ID)
	require.Error(t, err)
	require.Equal(t, err, pgx.ErrNoRows)
	require.Empty(t, r)
}