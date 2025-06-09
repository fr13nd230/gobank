package cache

import (
	"strconv"

	"github.com/fr13nd230/gobank/config"
	"github.com/redis/go-redis/v9"
)

func NewClient() *redis.Client {
    conf := newRedisOpts()
    return redis.NewClient(conf)
}

func newRedisOpts() *redis.Options {
    config.LoadConfig("../../.env")
    db, _ := strconv.ParseInt(config.GetVar("REDIS_DB"), 10, 64)
    return &redis.Options{
        Addr: config.GetVar("REDIS_ADDR"),
        DB: int(db),
    }
}