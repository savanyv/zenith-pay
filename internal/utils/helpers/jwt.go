package helpers

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/savanyv/zenith-pay/config"
)

const (
	jwtIssuer = "zenith-pay-backend"
	jwtAudiencePOS = "zenith-pay-pos"
	jwtAccessExpiry = 30 * time.Minute
)

type JWTClaims struct {
	UserID string `json:"user_id"`
	Username string `json:"username"`
	Role string `json:"role"`
	TokenVersion int `json:"token_version"`
	jwt.RegisteredClaims
}

type JWTService interface {
	GenerateAccessToken(userID, username, role string, tokenVersion int) (string, error)
	ValidateAccessToken(tokenString string) (*JWTClaims, error)
}

type jwtService struct {
	secretKey []byte
}

func NewJWTService() JWTService {
	secret := config.LoadConfig().JwtSecretKey
	if secret == "" {
		panic("JWT_SECRET_KEY is not set")
	}

	return &jwtService{
		secretKey: []byte(secret),
	}
}

func (j *jwtService) GenerateAccessToken(userID, username, role string, tokenVersion int) (string, error) {
	if userID == "" || username == "" || role == "" {
		return "", errors.New("invalid jwt payload")
	}

	claims := &JWTClaims{
		UserID: userID,
		Username: username,
		Role: role,
		TokenVersion: tokenVersion,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer: jwtIssuer,
			Audience: []string{jwtAudiencePOS},
			Subject: userID,
			IssuedAt: jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(jwtAccessExpiry)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.secretKey)
}

func (j *jwtService) ValidateAccessToken(tokenString string) (*JWTClaims, error) {
	token, err := jwt.ParseWithClaims(
		tokenString,
		&JWTClaims{},
		func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrInvalidKey
			}
			return j.secretKey, nil
		},
		jwt.WithIssuer(jwtIssuer),
		jwt.WithAudience(jwtAudiencePOS),
	)
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*JWTClaims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}
