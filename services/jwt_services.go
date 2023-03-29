package services

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type jwtService struct {
	secretkey string
	issure    string //quem está assiando o nosso token
}

func NewJwtService() *jwtService {
	return &jwtService{
		secretkey: "123456",
		issure:    "gin=api",
	}
}

type Claim struct {
	Sum uint `json:"sum"`
	jwt.StandardClaims
}

// funcao de gerar token // retorna uma string que será o nosso token ou um error
func (s *jwtService) GenerateToken(id uint) (string, error) {
	claim := &Claim{
		id,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 2).Unix(),
			Issuer:    s.issure,
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	t, err := token.SignedString([]byte(s.secretkey))
	if err != nil {
		return "", err
	}
	return t, nil
}

// validar o token

func (s *jwtService) Validatetoken(token string) bool {
	_, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, isValid := t.Method.(*jwt.SigningMethodHMAC); !isValid {
			return nil, fmt.Errorf("Invalid token: %v", token)
		}
		return []byte(s.secretkey), nil
	})
	return err == nil
}
