package accounts

import (
	"context"
	"fmt"
	"strconv"
	"time"

	rp "github.com/fr13nd230/gobank/database/repository"
	"github.com/fr13nd230/gobank/src/cache"
	"github.com/goccy/go-json"
	"github.com/jackc/pgx/v5/pgtype"
)

func CreateAccountProvider(
	ctx context.Context, 
	arg rp.NewAccountParams,
	q *rp.Queries,
) (rp.Account, error) {
	acc, err := q.NewAccount(ctx, arg)
	if err != nil {
		return rp.Account{}, err
	}
	return acc, nil
}

func ListAccountsProvider(
	ctx context.Context, 
	arg rp.ListAccountsParams,
	q *rp.Queries,
) ([]rp.Account, error) {
    var res []rp.Account
    cacheKey := "accounts:"+strconv.Itoa(int(arg.Offset))+":"+strconv.Itoa(int(arg.Limit))
    client := cache.NewClient()
    defer client.Close()
    
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

func FindAccountByIdProvider(
    ctx context.Context,
    id pgtype.UUID,
    q *rp.Queries,
) (rp.Account, error) {
    var res rp.Account
    cacheKey := fmt.Sprintf("account:%v", id)
    client := cache.NewClient()
    defer client.Close()
    
    val, _ := client.Get(ctx, cacheKey).Result()
    if len(val) > 0 {
        if err := json.Unmarshal([]byte(val), &res); err != nil {
            return rp.Account{}, err
        }
        
        return res, nil
    }
    
    acc, err := q.FindAccountById(ctx, id)
    if err != nil {
        return rp.Account{}, err
    }
    
    data, err := json.Marshal(acc)
    if err != nil {
        return rp.Account{}, nil
    }
    
    if err := client.Set(ctx, cacheKey, data, 15*time.Minute).Err(); err != nil {
       return rp.Account{}, err
   }

   return acc, nil
}