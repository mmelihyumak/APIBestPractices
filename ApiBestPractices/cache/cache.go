package cache

import (
    "context"
    "time"
    "github.com/redis/go-redis/v9"
    "log"
)

var ctx = context.Background()

type RedisCache struct {
    client *redis.Client
}

// Redis bağlantısını kuran yapılandırıcı
func NewRedisCache(addr, password string, db int) *RedisCache {
    rdb := redis.NewClient(&redis.Options{
        Addr:     addr,
        Password: password,
        DB:       db,
    })

    log.Println("Redis bağlantısı başarılı")

    return &RedisCache{
        client: rdb,
    }
}

// Cache'e veri yazma
func (r *RedisCache) Set(key string, value interface{}, expiration time.Duration) error {
    log.Println("Başarılı şekilde redise yazıldı")
    return r.client.Set(ctx, key, value, expiration).Err()
}

// Cache'ten veri okuma
func (r *RedisCache) Get(key string) (string, error) {
    log.Println("Başarılı şekilde redisten okundu")
    return r.client.Get(ctx, key).Result()
}

// Cache'teki veriyi silme
func (r *RedisCache) Delete(key string) error {
    log.Println("Başarılı şekilde redisten silindi")
    return r.client.Del(ctx, key).Err()
}
