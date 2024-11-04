package services

import (
    "encoding/json"
    "time"
    "goapibestpractices/cache"
    "goapibestpractices/repository"
    "fmt"
)

type UserService struct {
    repo  *repository.UserRepository
    cache *cache.RedisCache
}

func NewUserService(repo *repository.UserRepository, cache *cache.RedisCache) *UserService {
    return &UserService{
        repo:  repo,
        cache: cache,
    }
}

// Cache kontrolü ile kullanıcıyı alır
func (s *UserService) GetUserByID(id int64) (*repository.User, error) {
    cacheKey := fmt.Sprintf("user:%d", id)
    
    // Cache'ten kontrol et
    cachedUser, err := s.cache.Get(cacheKey)
    if err == nil && cachedUser != "" {
        var user repository.User
        if err := json.Unmarshal([]byte(cachedUser), &user); err == nil {
            return &user, nil
        }
    }
    
    // Cache'te yoksa veritabanından getir ve cache'e yaz
    user, err := s.repo.GetUserByID(id)
    if err != nil {
        return nil, err
    }

    userJSON, err := json.Marshal(user)
    if err == nil {
        _ = s.cache.Set(cacheKey, string(userJSON), 5*time.Minute) // 5 dakika cache'te tut
    }
    
    return user, nil
}

// Kullanıcı oluştururken cache'i güncelle
func (s *UserService) CreateUser(user *repository.User) error {
    if err := s.repo.CreateUser(user); err != nil {
        return err
    }

    // Yeni kullanıcıyı cache'e ekle
    cacheKey := fmt.Sprintf("user:%d", user.ID)
    userJSON, _ := json.Marshal(user)
    _ = s.cache.Set(cacheKey, string(userJSON), 5*time.Minute)

    return nil
}

// Kullanıcıyı güncellerken cache'i de güncelle
func (s *UserService) UpdateUser(user *repository.User) error {
    if err := s.repo.UpdateUser(user); err != nil {
        return err
    }

    // Cache'i güncelle
    cacheKey := fmt.Sprintf("user:%d", user.ID)
    userJSON, _ := json.Marshal(user)
    _ = s.cache.Set(cacheKey, string(userJSON), 5*time.Minute)

    return nil
}

// Kullanıcı silinirken cache'i de sil
func (s *UserService) DeleteUser(id int64) error {
    if err := s.repo.DeleteUser(id); err != nil {
        return err
    }

    // Cache'ten kaldır
    cacheKey := fmt.Sprintf("user:%d", id)
    _ = s.cache.Delete(cacheKey)

    return nil
}
