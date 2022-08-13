package auth

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha1"
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/kimxuanhong/go-campaign-no-02/pkg/dto"
	"github.com/labstack/echo/v4/middleware"
	"io"
	"os"
	"time"
)

type CustomClaims struct {
	User dto.User `json:"user"`
	jwt.StandardClaims
}

type JwtConfig interface {
	GetJWTSecret() string
	GenerateAccessToken(user *dto.User) (dto.Token, error)
	GenerateRefreshToken(model dto.Token) (dto.Token, error)
	ValidateToken(accessToken string) (dto.User, error)
	ValidateRefreshToken(token dto.Token) (dto.User, error)
	GenerateToken(user *dto.User, expirationTime time.Time, secret []byte) (dto.Token, error)
	GetConfig() middleware.JWTConfig
}

type JwtConfigImpl struct {
}

var instanceJwtConfig *JwtConfigImpl

func NewJwtConfig() *JwtConfigImpl {
	if instanceJwtConfig == nil {
		instanceJwtConfig = &JwtConfigImpl{}
	}
	return instanceJwtConfig
}

func (r *JwtConfigImpl) GetJWTSecret() string {
	return os.Getenv("JWT_SECRET_KEY")
}

func (r *JwtConfigImpl) GenerateAccessToken(user *dto.User) (dto.Token, error) {
	// Declare the expiration time of the token (1h).
	expirationTime := time.Now().Add(1 * time.Hour)

	return r.GenerateToken(user, expirationTime, []byte(r.GetJWTSecret()))
}

func (r *JwtConfigImpl) GenerateRefreshToken(model dto.Token) (dto.Token, error) {
	hash := sha1.New()
	_, err := io.WriteString(hash, r.GetJWTSecret())
	if err != nil {
		return model, err
	}

	salt := string(hash.Sum(nil))[0:16]
	block, err := aes.NewCipher([]byte(salt))
	if err != nil {
		fmt.Println(err.Error())

		return model, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return model, err
	}

	nonce := make([]byte, gcm.NonceSize())
	_, err = io.ReadFull(rand.Reader, nonce)
	if err != nil {
		return model, err
	}

	model.RefreshToken = base64.URLEncoding.EncodeToString(gcm.Seal(nonce, nonce, []byte(model.AccessToken), nil))

	return model, nil

}

func (r *JwtConfigImpl) GenerateToken(user *dto.User, expirationTime time.Time, secret []byte) (dto.Token, error) {
	// Create the JWT claims, which includes the username and expiry time.
	claims := &CustomClaims{
		User: *user,
		StandardClaims: jwt.StandardClaims{
			// In JWT, the expiry time is expressed as unix milliseconds.
			ExpiresAt: expirationTime.Unix(),
		},
	}

	// Declare the token with the HS256 algorithm used for signing, and the claims.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Create the JWT string.
	jwtToken := dto.Token{}
	tokenString, err := token.SignedString(secret)
	if err != nil {
		return jwtToken, err
	}

	jwtToken.AccessToken = tokenString
	return r.GenerateRefreshToken(jwtToken)
}

func (r *JwtConfigImpl) ValidateToken(accessToken string) (dto.User, error) {
	token, err := jwt.ParseWithClaims(accessToken, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(r.GetJWTSecret()), nil
	})

	user := dto.User{}
	if err != nil {
		return user, err
	}

	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims.User, nil
	}

	return user, errors.New("invalid token")
}

func (r *JwtConfigImpl) ValidateRefreshToken(model dto.Token) (dto.User, error) {
	hash := sha1.New()
	_, err := io.WriteString(hash, r.GetJWTSecret())
	if err != nil {
		return dto.User{}, err
	}

	user := dto.User{}
	salt := string(hash.Sum(nil))[0:16]
	block, err := aes.NewCipher([]byte(salt))
	if err != nil {
		return user, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return user, err
	}

	data, err := base64.URLEncoding.DecodeString(model.RefreshToken)
	if err != nil {
		return user, err
	}

	nonceSize := gcm.NonceSize()
	nonce, ciphertext := data[:nonceSize], data[nonceSize:]

	plain, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return user, err
	}

	if string(plain) != model.AccessToken {
		return user, errors.New("invalid token")
	}

	claims := &CustomClaims{}
	parser := jwt.Parser{}
	token, _, err := parser.ParseUnverified(model.AccessToken, claims)

	if claims, ok := token.Claims.(*CustomClaims); ok {
		user = claims.User
		return user, nil
	}

	return user, errors.New("invalid token")

}

func (r *JwtConfigImpl) GetConfig() middleware.JWTConfig {
	config := middleware.JWTConfig{
		Claims:     &CustomClaims{},
		SigningKey: []byte(r.GetJWTSecret()),
	}

	return config
}
