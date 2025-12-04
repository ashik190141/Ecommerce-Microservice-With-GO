package service

import (
	"Product-Service/dto"
	"Product-Service/interfaces"
	"context"
	"encoding/json"
	"log"
	"time"
	"github.com/redis/go-redis/v9"
)

type RedisService struct {
	redis *redis.Client
	ctx   context.Context
}

func NewRedisService(redisClient *redis.Client) interfaces.ProductRedis {
	return &RedisService{
		redis: redisClient,
		ctx:   context.Background(),
	}
}

func (r *RedisService) GetProductFromCache(key string) ([]dto.GetProductResponse, error) {
	val, err := r.redis.HGetAll(r.ctx, key).Result()
	if err != nil {
		return []dto.GetProductResponse{}, err
	}

	var products []dto.GetProductResponse

	for _, data := range val {
		var p dto.GetProductResponse
		err := json.Unmarshal([]byte(data), &p)
		if err != nil {
			log.Println("Failed to unmarshal product:", err)
			continue
		}
		products = append(products, p)
	}

	remainingTTL, _ := r.redis.TTL(r.ctx, key).Result()
	log.Printf("Remaining TTL for '%s': %v\n", key, remainingTTL)

	r.SetExpireTimeFromCache(key, 5*60*time.Second)
	againRemainingTTL, _ := r.redis.TTL(r.ctx, key).Result()
	log.Printf("Set Again Remaining TTL for '%s': %v\n", key, againRemainingTTL)

	return products, nil
}

func (r *RedisService) SetProductToCache(key string, product dto.GetProductResponse) bool {
	data, err := json.Marshal(product)
	if err != nil {
		return false
	}

	res1, err := r.redis.HSet(r.ctx, key, product.Sku, data).Result()
	if err != nil || res1 == 0 {
		log.Println("Failed to cache product:", err)
		return false
	}

	r.SetExpireTimeFromCache(key, 5*60*time.Second)

	return true
}

func (r *RedisService) SetExpireTimeFromCache(key string, duration time.Duration) {
	err := r.redis.Expire(r.ctx, key, duration).Err()
	if err != nil {
		log.Println("Failed to set TTL:", err)
	}
}