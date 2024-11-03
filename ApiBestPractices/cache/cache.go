package cache

import (
    "context"
    "github.com/redis/go-redis/v9"
    "sync"
    "time"
    "log"
)

// Cache interface'i, cache işlemlerini soyutlamak için kullanılır
type Cache interface {
    Set(key string, value interface{}, expiration time.Duration) error
    Get(key string) (string, error)
    Delete(key string) error
}

// RedisCache yapısı, Redis bağlantısını içerir
type RedisCache struct {
    client *redis.Client
}

// Redis cache için yeni bir RedisCache nesnesi oluşturur
func NewRedisCache(client *redis.Client) Cache {
    return &RedisCache{client: client}
}

// Set, cache'e veri ekler
func (r *RedisCache) Set(key string, value interface{}, expiration time.Duration) error {
    log.Print("Cache'e veri eklendi")
    return r.client.Set(context.Background(), key, value, expiration).Err()
}

// Get, cache'ten veri çeker
func (r *RedisCache) Get(key string) (string, error) {
    log.Print("Cache'ten veri alındı")
    return r.client.Get(context.Background(), key).Result()
}

// Delete, cache'ten veriyi siler
func (r *RedisCache) Delete(key string) error {
    log.Print("Cache'ten veri silindi")
    return r.client.Del(context.Background(), key).Err()
}

// Redis bağlantısını singleton olarak kurar
var (
    once sync.Once
    rdb  *redis.Client
)

// GetRedisClient, tek bir Redis bağlantısı sağlar
func GetRedisClient() *redis.Client {
    once.Do(func() {
        rdb = redis.NewClient(&redis.Options{
            Addr:     "localhost:6379", // Redis sunucu adresi
            Password: "",               // Redis şifresi yoksa boş bırakın
            DB:       0,                // Kullanılacak Redis veritabanı
        })
    })
    return rdb
}
