package repository

import (
	"context"
	"testing"

	"github.com/fr13nd230/gobank/utils"
	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/require"
)

// CreateAccount will create for use new accounts to be used
// in future unit tests, this way is more cleaner.
func CreateAccount (t *testing.T) Account {
	accountParams := NewAccountParams {
		Owner: utils.GenRandName(8),
		Currency: utils.GenRandCurrency("USD", "EUR"),
	}
	
	acc, err := queries.NewAccount(context.Background(), accountParams)
	
	require.NoError(t, err)
	require.NotEmpty(t, acc)
	require.Equal(t, accountParams.Owner, acc.Owner)
	require.Equal(t, accountParams.Currency, acc.Currency)
	require.NotZero(t, acc.ID)
	require.Zero(t, acc.Balance)
	require.NotZero(t, acc.CreatedAt)
	require.NotZero(t, acc.UpdatedAt)

	return acc
}

// TestCreateAccount will test the generated SQLC code for
// adding new account, no mocking is introduced in here maybe later.
func TestCreateAccount(t *testing.T) {
	CreateAccount(t)
}

// TestListAccounts will test the accounts retrievals from DB then 
// assert each individual object.
func TestListAccounts (t *testing.T) {
	for i:=0; i < 10; i++ {
		CreateAccount(t)
	}
	
	limit := int32(10)
	offset := int32(1)
	arg := ListAccountsParams{
		Limit: limit,
		Offset: offset,
	}
	
	accs, err := queries.ListAccounts(context.Background(), arg)
	
	require.NoError(t, err)
	require.NotEmpty(t, accs)
	require.LessOrEqual(t, len(accs), int(limit))
	
	for _, acc := range accs{
		require.NotEmpty(t, acc)
		require.NotZero(t, acc.ID)
		require.NotZero(t, acc.Owner)
		require.Zero(t, acc.Balance)
		require.NotZero(t, acc.CreatedAt)
		require.NotZero(t, acc.UpdatedAt)
	}
}

// TestFindAccountById this will test account retrieval by creating new account
// then try to retrieve if from the database.
func TestFindAccountById(t *testing.T) {
	acc := CreateAccount(t)
	
	trgAcc, err := queries.FindAccountById(context.Background(), acc.ID)
	
	require.NoError(t, err)
	require.NotEmpty(t, trgAcc)
	require.Equal(t, acc, trgAcc)
}

// TestUpdateAccountById this will test account information update by creating 
// new account then updates it.
func TestUpdateAccountById(t *testing.T) {
	oldAcc := CreateAccount(t)

	arg := UpdateAccountByIdParams{
		ID: oldAcc.ID,
		Owner: utils.GenRandName(10),
		Currency: utils.GenRandCurrency("DZD", "CAD"),
	}
	newAcc, err := queries.UpdateAccountById(context.Background(), arg)
	
	require.NoError(t, err)
	require.NotEmpty(t, newAcc)
	require.NotZero(t, newAcc.Owner)
	require.NotZero(t, newAcc.Currency)
	require.NotEqual(t, newAcc.Owner, oldAcc.Owner)
	require.NotEqual(t, newAcc.Currency, oldAcc.Currency)
}

// TestDeleteAccountById this will attempt to create a whole new account
// then attempts to delete it using its ID.
func TestDeleteAccountById(t *testing.T) {
	acc := CreateAccount(t)
	
	err := queries.DeleteAccountById(context.Background(), acc.ID)
	require.NoError(t, err)
	
	// Check if the account still exists
	find, err := queries.FindAccountById(context.Background(), acc.ID)
	require.Error(t, err)
	require.Equal(t, err, pgx.ErrNoRows)
	require.Empty(t, find)
}
