package services

import (
	"fmt"
	"placement/models"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// JwtService interface defining what methods are available in this service
type JwtService interface {
	GenerateAccessToken(userID string, roll models.Roll) (string, error)
	VerifyAccessToken(token string) (*jwt.Token, error)

	GenerateRefreshToken(userID string, roll models.Roll) (string, error)
	VerifyRefreshToken(token string) (*jwt.Token, error)
}

type accessJwtClaim struct {
	UserID string      `json:"userID"`
	Roll   models.Roll `json:"roll"`
	jwt.StandardClaims
}
type refreshJwtClaim struct {
	UserID string      `json:"userID"`
	Roll   models.Roll `json:"roll"`
	jwt.StandardClaims
}
type jwtService struct {
	accessSecret  string
	refreshSecret string
	issuer        string
}

// NewJwtService creates a new JWT service
func NewJwtService() JwtService {
	return &jwtService{
		accessSecret:  "accessSecret",
		issuer:        "issuer",
		refreshSecret: "refreshSecret",
	}
}

// Generate generates a new JWT token
func (jwtService *jwtService) GenerateAccessToken(userID string, roll models.Roll) (string, error) {
	claim := &accessJwtClaim{
		userID,
		roll,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 30).Unix(),
			Issuer:    jwtService.issuer,
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	t, err := token.SignedString([]byte(jwtService.accessSecret))
	if err != nil {
		return "", err
	}
	return t, nil
}

func (jwtService *jwtService) GenerateRefreshToken(userID string, roll models.Roll) (string, error) {
	claim := &refreshJwtClaim{
		userID,
		roll,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24 * 30).Unix(),
			Issuer:    jwtService.issuer,
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	t, err := token.SignedString([]byte(jwtService.refreshSecret))
	if err != nil {
		return "", err
	}
	return t, nil
}

// Verify verifies the jet token
func (jwtService *jwtService) VerifyAccessToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(jwtService.accessSecret), nil
	})
}

func (jwtService *jwtService) VerifyRefreshToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(jwtService.refreshSecret), nil
	})
}
