package accounts

import (
	"context"
	"strconv"
	"time"

	"github.com/fr13nd230/gobank/database/repository"
	"github.com/fr13nd230/gobank/src/cache"
	"github.com/goccy/go-json"
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

func ListAccountsProviderfunc(
	ctx context.Context, 
	arg repository.ListAccountsParams,
	q *repository.Queries,
) ([]repository.Account, error) {
    var res []repository.Account
    cacheKey := "accounts:"+strconv.Itoa(int(arg.Offset))+":"+strconv.Itoa(int(arg.Limit))
    client := cache.NewClient()
    
    val, _ := client.Get(ctx, cacheKey).Result()
    if len(val) > 0 {
        if err := json.Unmarshal([]byte(val), &res); err != nil {
            return nil, err
        }
        
        return res, nil
    }

    accs, err := q.ListAccounts(ctx, arg)
    if err != nil {
        return nil, err
    }
    
    data, err := json.Marshal(accs)
    if err != nil {
        return nil, err
    }
    if err := client.Set(ctx, cacheKey, data, 30*time.Minute).Err(); err != nil {
        return nil, err
    }
    
    return accs, nil
}