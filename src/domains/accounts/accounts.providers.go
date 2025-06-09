package accounts

import (
	"context"

	"github.com/fr13nd230/gobank/database/repository"
)

func CreateAccountProvider(
	ctx context.Context, 
	arg repository.NewAccountParams,
	q *repository.Queries,
) (repository.Account, error) {
	acc, err := q.NewAccount(ctx, arg)
	if err != nil {
		return repository.Account{}, err
	}
	return acc, nil
}