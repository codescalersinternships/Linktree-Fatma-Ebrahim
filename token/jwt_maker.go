package token

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type JWTMaker struct {
	secretKey string
}

type UserClaims struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Username string             `bson:"username"`
	Email    string             `bson:"email"`
	Password string             `bson:"password"`
	jwt.RegisteredClaims
}

func NewUserClaims(id primitive.ObjectID, username string, email string, password string, duration time.Duration) (*UserClaims, error) {
	tokenID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}
	return &UserClaims{
		ID:       id,
		Username: username,
		Email:    email,
		Password: password,
		RegisteredClaims: jwt.RegisteredClaims{
			ID:        tokenID.String(),
			Subject:   username, //unique
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(duration)),
		},
	}, nil
}

func NewJWTMaker(secretKey string) *JWTMaker {
	return &JWTMaker{secretKey}
}

func (maker *JWTMaker) CreateToken(id primitive.ObjectID, username string, email string, password string, duration time.Duration) (string, *UserClaims, error) {
	claims, err := NewUserClaims(id, username, email, password, duration)
	if err != nil {
		return "", nil, err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(maker.secretKey))
	if err != nil {
		return "", nil, fmt.Errorf("could not sign token: %v", err)
	}
	return tokenString, claims, nil

}

func (maker *JWTMaker) VerifyToken(tokenStr string) (*UserClaims, error) {
	token,err:=jwt.ParseWithClaims(tokenStr, &UserClaims{},func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, fmt.Errorf("invalid token signing method")
		}

		return []byte(maker.secretKey), nil
	})
	if err != nil {
		return nil, fmt.Errorf("error parsing token: %v", err)
	}

	claims, ok := token.Claims.(*UserClaims)	
	if !ok {
		return nil, fmt.Errorf("invalid token claims")
		
	}
	return claims, nil
	
}
