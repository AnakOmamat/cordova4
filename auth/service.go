package auth

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type Service interface {
	GenerateToken(userID int) (string, error)
	ValidateToken(Token string) (*jwt.Token, error)
}

type jwtService struct {
}

var SECRET_KEY = []byte("BOGIECORDOVA4_s3cr3T_k3Y")

func NewService() *jwtService {
	return &jwtService{}

}

func (s *jwtService) GenerateToken(userID int) (string, error) {
	//claim := jwt.MapClaims{}
	//claim["user_id"] = userID

	expirationTime := time.Now().Add(1 * time.Minute)

	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     expirationTime.Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString(SECRET_KEY)
	if err != nil {
		return signedToken, err
	}

	return signedToken, nil

}

func (s *jwtService) ValidateToken(encodedToken string) (*jwt.Token, error) {
	token, err := jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, errors.New("Invalid Token")
		}
		return []byte(SECRET_KEY), nil

	})
	if err != nil {
		return token, err
	}
	return token, nil
}
