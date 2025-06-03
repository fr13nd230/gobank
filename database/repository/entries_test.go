package repository

import (
	"context"
	"testing"

	"github.com/fr13nd230/gobank/utils"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/require"
)

// CreateEntry will test the creation of a new entry record
// it requires CreateTransfer to work.
func CreateEntry(t *testing.T) (Entry, Entry, Transfer) {
	tr := CreateTrasnfer(t)

	arg1 := NewEntryParams{
		AccountID: tr.FromAcc,
		Amount: -tr.Amount,
		TransferID: tr.ID,
	}
	arg2 := NewEntryParams{
		AccountID: tr.ToAcc,
		Amount: tr.Amount,
		TransferID: tr.ID,
	}
	
	en1, err1 := queries.NewEntry(context.Background(), arg1)
	en2, err2 := queries.NewEntry(context.Background(), arg2)
	require.NoError(t, err1)
	require.NoError(t, err2)
	require.NotEmpty(t, en1)
	require.NotEmpty(t, en2)
	
	require.NotZero(t, en1.ID)
	require.NotZero(t, en1.AccountID)
	require.NotZero(t, en1.TransferID)
	require.NotZero(t, en1.Amount)
	require.Less(t, en1.Amount, tr.Amount)
	require.NotZero(t, en1.CreatedAt)
	require.NotZero(t, en1.UpdatedAt)
	
	require.NotZero(t, en2.ID)
	require.NotZero(t, en2.AccountID)
	require.NotZero(t, en2.TransferID)
	require.NotZero(t, en2.Amount)
	require.Equal(t, en2.Amount, tr.Amount)
	require.NotZero(t, en2.CreatedAt)
	require.NotZero(t, en2.UpdatedAt)
	
	return en1, en2, tr
}

// TestCreateEntry will test if create entry query function
// is valid and do its purpose.
func TestCreateEntry(t *testing.T) {
	CreateEntry(t)
}

// TestListEntries will test if list entries query function
// does function as expected.
func TestListEntries(t *testing.T) {
	limit := int32(10)
	offset := int32(1)
	// Create 10 new Entries
	for i:=0; i<int(limit); i++{
		CreateEntry(t)
	}
	
	arg := ListEntriesParams{
		Limit: limit,
		Offset: offset,
	}
	
	ens, err := queries.ListEntries(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, ens)
	require.LessOrEqual(t, int(limit), len(ens))
	
	for _, en := range ens {
		require.NotEmpty(t, en)
		require.NotZero(t, en.ID)
		require.NotZero(t, en.AccountID)
		require.NotZero(t, en.TransferID)
		require.NotZero(t, en.Amount)
		require.NotZero(t, en.CreatedAt)
		require.NotZero(t, en.UpdatedAt)
	}
}

// TestFindEntryById will test if find entry by its ID
// does work as expected, requires CreateEntry to work.
func TestFindEntryById(t *testing.T){
    en1, en2, _ := CreateEntry(t)

    entries := []pgtype.UUID{en1.ID, en2.ID}

    for _, id := range entries {
        find, err := queries.FindEntryById(context.Background(), id)

        require.NoError(t, err)
        require.NotEmpty(t, find)
        require.NotZero(t, find.ID)
        require.Equal(t, id, find.ID)
        require.NotZero(t, find.AccountID)
        require.NotZero(t, find.TransferID)
        require.NotZero(t, find.Amount)
        require.NotZero(t, find.CreatedAt)
        require.NotZero(t, find.UpdatedAt)
    }
}

// TestFindEntriesByTrAcc this will test find entries queries
// parameters and see if it does handle fetching by transfer_id and account_id.
// func TestFindEntriesByTrAcc(t *testing.T){
// 	_, _, tr := CreateEntry(t)
// 	limit := int32(10)
// 	offset := int32(1)
	
// 	args := []FindEntriesByTrAccParams{
// 		{
// 			AccountID: tr.FromAcc,
// 			Limit: limit,
// 			Offset: offset,
// 		},
// 		{
// 			AccountID: tr.ToAcc,
// 			Limit: limit,
// 			Offset: offset,
// 		},
// 		{
// 			AccountID: tr.ID,
// 			Limit: limit,
// 			Offset: offset,
// 		},
// 	}
	
// 	for _, arg := range args {
// 		ens, err := queries.FindEntriesByTrAcc(context.Background(), arg)	
// 		require.NoError(t, err)
//         require.NotEmpty(t, ens)
//         require.LessOrEqual(t, int(limit), len(ens))
        
//         for _, en := range ens {
//         	require.NotZero(t, en.ID)
//         	require.NotZero(t, en.AccountID)
//         	require.NotZero(t, en.TransferID)
//         	require.NotZero(t, en.Amount)
//         	require.NotZero(t, en.CreatedAt)
//         	require.NotZero(t, en.UpdatedAt)
//         }
// 	}
// }

// TestUpdateEntryById will test the update entry function
// query and validate its functionality.
func TestUpdateEntryById(t *testing.T){
	en1, _, _ := CreateEntry(t)
	
	arg := UpdateEntryByIdParams{
		ID: en1.ID,
		Amount: utils.GenRandMoney(10, 15),
	}
	
	newEn, err := queries.UpdateEntryById(context.Background(), arg)
	
	require.NoError(t, err)
	require.NotEmpty(t, newEn)
	require.NotZero(t, newEn)
	require.NotEqual(t, en1.Amount, newEn)
}

// TestDeleteEntryById will test the delete entry function
// query and validate its functionality.
func TestDeleteEntryById(t* testing.T){
	en, _, _ := CreateEntry(t)
	
	err := queries.DeleteEntryById(context.Background(), en.ID)
	require.NoError(t, err)
	
	find, err := queries.FindEntryById(context.Background(), en.ID)
	require.Error(t, err)
	require.Equal(t, err, pgx.ErrNoRows)
	require.Empty(t, find)
}
