package helpers

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/savanyv/zenith-pay/config"
)

const (
	expectedSecretHash = "d4637482715b4d6efc627fd7b1310e711de536810d5121954b0463672399868f"
	jwtIssuer = "zenith-pay-backend"
	jwtExpiry = 24 * time.Hour
)

type JWTClaims struct {
	UserID string `json:"user_id"`
	Username string `json:"username"`
	Role string `json:"role"`
	jwt.RegisteredClaims
}

type JWTService interface {
	GenerateToken(userID, username, role string) (string, error)
	ValidateToken(tokenString string) (*JWTClaims, error)
}

type jwtService struct {
	secretKey string
	issuer string
}

func NewJWTService() JWTService {
	secret := config.LoadConfig().JwtSecretKey

	hash := sha256.Sum256([]byte(secret))
	if hex.EncodeToString(hash[:]) != expectedSecretHash {
		log.Fatal("JWT_SECRET_KEY is invalid or has been changed!")
	}

	decode, err := base64.StdEncoding.DecodeString(secret)
	if err != nil {
		log.Fatal("JWT_SECRET_KEY is invalid or has been changed!")
	}

	return &jwtService{
		secretKey: string(decode),
		issuer: jwtIssuer,
	}
}

func (j *jwtService) GenerateToken(userID, username, role string) (string, error) {
	if userID == "" || username == "" || role == "" {
		return "", errors.New("invalid token payload")
	}

	claims := &JWTClaims{
		UserID: userID,
		Username: username,
		Role: role,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer: j.issuer,
			IssuedAt: jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(jwtExpiry)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(j.secretKey))
}

func (j *jwtService) ValidateToken(tokenString string) (*JWTClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrInvalidKey
		}
		return []byte(j.secretKey), nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*JWTClaims)
	if !ok || !token.Valid {
		return nil, jwt.ErrInvalidKey
	}

	if claims.Issuer != j.issuer {
		return nil, jwt.ErrInvalidKey
	}

	return claims, nil
}
