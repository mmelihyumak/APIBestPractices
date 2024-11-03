package services

import (
    "time"
    "goapibestpractices/cache"
)

// UserService yapısı, kullanıcı işlemlerini tanımlar
type UserService struct {
    cache cache.Cache
}

// Yeni bir UserService oluşturur
func NewUserService(cache cache.Cache) *UserService {
    return &UserService{cache: cache}
}

// GetUser, kullanıcı verisini cache'ten veya kaynaktan alır
func (s *UserService) GetUser(id string) (string, error) {
    // İlk olarak cache'ten veriyi çekmeye çalış
    cachedData, err := s.cache.Get(id)
    if err == nil {
        return cachedData, nil
    }

    // Veri cache'te bulunamazsa kaynaktan veriyi getir (örneğin veritabanından)
    dataFromDB := "userDataFromDB" // Bu kısım gerçek veritabanı sorgusu ile değiştirilmeli

    // Veriyi cache'e kaydet
    _ = s.cache.Set(id, dataFromDB, 10*time.Minute)
    return dataFromDB, nil
}
