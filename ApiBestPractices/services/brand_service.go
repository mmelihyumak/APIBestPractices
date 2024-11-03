package services

import (
	"encoding/json"
	"goapibestpractices/cache"
	"goapibestpractices/repository"
	"time"
)

// UserService yapısı, kullanıcı işlemlerini tanımlar
type BrandService struct {
    cache cache.Cache
}

// Yeni bir UserService oluşturur
func NewBrandService(cache cache.Cache) *BrandService {
    return &BrandService{cache: cache}
}

func (s *BrandService) GetBrand(id string) (string, error) {
    // İlk olarak cache'ten veriyi çekmeye çalış
    cachedData, err := s.cache.Get(id)
    if err == nil {
        return cachedData, nil
    }

    // Veri cache'te bulunamazsa kaynaktan veriyi getir (örneğin veritabanından)
    dataFromDB, err := repository.GetAllBrands() // Bu kısım gerçek veritabanı sorgusu ile değiştirilmeli
    if err != nil{
        return "", err
    }
    
    jsonData, err := json.Marshal(dataFromDB)
    if err != nil{
        return "", err
    }

    // Veriyi cache'e kaydet
    _ = s.cache.Set(id, dataFromDB, 10*time.Minute)
    return string(jsonData), nil
}