package repository

import (
	"address-book-go/internal/token"
	"github.com/go-redis/redis/v7"
	"strconv"
)

type Cache struct {
	redisCluster *redis.ClusterClient
}

func NewRedis(rc *redis.ClusterClient) token.Cache {
	return &Cache{
		redisCluster: rc,
	}
}

func (c *Cache) GetTokenIat(accId int) (float64, error) {
	key := token.GetSysAccountTokenRedisKey(accId)
	iatStr, err := c.redisCluster.HGet(key, "iat").Result()
	if err != nil {
		return 0, err
	}

	iat, err := strconv.ParseFloat(iatStr, 64)

	return iat, err
}

func (c *Cache) SetTokenIat(accId int, iat int64) error {
	key := token.GetSysAccountTokenRedisKey(accId)
	err := c.redisCluster.HSet(key, "iat", iat).Err()
	return err
}
