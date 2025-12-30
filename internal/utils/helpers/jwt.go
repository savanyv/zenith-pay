package helpers

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/savanyv/zenith-pay/config"
)

type JWTClaim struct {
	UserID string `json:"user_id"`
	Username string `json:"username"`
	Role string `json:"role"`
	jwt.RegisteredClaims
}

type JWTService interface {
	GenerateToken(userID, username, role string) (string, error)
	ValidateToken(tokenString string) (*JWTClaim, error)
}

type jwtService struct {
	secretKey string
	issuer string
}

func NewJWTService() JWTService {
	return &jwtService{
		secretKey: config.LoadConfig().JwtSecretKey,
		issuer: "zenith-pay-backend",
	}
}

func (j *jwtService) GenerateToken(userID, username, role string) (string, error) {
	claims := &JWTClaim{
		UserID: userID,
		Username: username,
		Role: role,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer: "zenith-pay-backend",
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt: jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(j.secretKey))
}

func (j *jwtService) ValidateToken(tokenString string) (*JWTClaim, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JWTClaim{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(j.secretKey), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*JWTClaim)
	if !ok || !token.Valid {
		return nil, jwt.ErrInvalidKey
	}

	return claims, nil
}
