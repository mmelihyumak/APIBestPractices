package middlewares

import (
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"strings"
	"fmt"
	"time"
)

var secretKey = []byte("your_secret_key")

func VerifyToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// İmza yönteminin HS256 olup olmadığını kontrol et
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secretKey, nil
	})
}

func CreateToken(userID string) (string, error) {
    claims := jwt.MapClaims{
        "authorized": true,
        "user_id":    userID,
        "exp":        time.Now().Add(time.Hour * 24).Unix(),
    }
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString([]byte("your_secret_key"))
}

func JwtAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Authorization header'ını al
		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			http.Error(w, "Yetkisiz erişim: Token gerekli", http.StatusUnauthorized)
			return
		}

		// "Bearer " kısmını çıkart
		tokenString = strings.TrimPrefix(tokenString, "Bearer ")

		// Token doğrula
		token, err := VerifyToken(tokenString)
		if err != nil || !token.Valid {
			http.Error(w, "Yetkisiz erişim: Geçersiz token", http.StatusUnauthorized)
			return
		}

		// Devam et
		next.ServeHTTP(w, r)
	})
}
