package utils

import (
	"MengGoods/config"
	"MengGoods/pkg/constants"
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

type Claims struct {
	Type int64 `json:"type"`
	Uid  int64 `json:"uid"`
	jwt.StandardClaims
}

// 创建两种网关token,一种是access token，第二种是 refresh token
func CreateGatewayToken(uid int64) (string, string, error) {
	accessToken, err := CreateToken(constants.TypeAccess, uid) //创建access token
	if err != nil {
		return "", "", err
	}
	refreshToken, err := CreateToken(constants.TypeRefresh, uid) //创建refresh token
	if err != nil {
		return "", "", err
	}
	return accessToken, refreshToken, nil
}

// 根据token的Type和用户uid创建token
func CreateToken(tokenType int64, uid int64) (string, error) {
	var expiredDurationStr string
	switch tokenType {
	case constants.TypeAccess:
		expiredDurationStr = config.Conf.JWT.AccessExpire
	case constants.TypeRefresh:
		expiredDurationStr = config.Conf.JWT.RefreshExpire
	case constants.TypeLogin:
		expiredDurationStr = config.Conf.JWT.AccessExpire
	}
	expiredDuration, err := time.ParseDuration(expiredDurationStr)
	if err != nil {
		return "", err
	}
	claims := Claims{
		Type: tokenType,
		Uid:  uid,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(expiredDuration).Unix(),
			Issuer:    config.Conf.JWT.Issuer,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodEdDSA, claims)
	privateKey, err := parsePrivateKey(config.Conf.JWT.PrivateKey)
	if err != nil {
		return "", err
	}
	tokenString, err := token.SignedString(privateKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// jwt解析私钥
func parsePrivateKey(key string) (interface{}, error) {
	parsedKey, err := jwt.ParseEdPrivateKeyFromPEM([]byte(key))
	if err != nil {
		return nil, err
	}
	return parsedKey, nil
}

// jwt解析公钥
func parsePublicKey(key string) (interface{}, error) {
	parsedKey, err := jwt.ParseEdPublicKeyFromPEM([]byte(key))
	if err != nil {
		return nil, err
	}
	return parsedKey, nil
}

// 验证并解析claims
func verifyToken(token string, key interface{}) (*Claims, error) {
	parsedToken, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodEd25519); !ok {
			return nil, fmt.Errorf("token method wrong, got %v", token.Method)
		}
		return key, nil
	})
	if err != nil {
		return nil, err
	}
	claims, ok := parsedToken.Claims.(*Claims)
	if !ok || !parsedToken.Valid {
		return nil, errors.New("token invalid")
	}
	return claims, nil
}

// 验证token
func CheckToken(token string) (*Claims, error) {
	publicKey, err := parsePublicKey(config.Conf.JWT.PublicKey)
	if err != nil {
		return nil, err
	}
	claims, err := verifyToken(token, publicKey)
	if err != nil {
		return nil, err
	}
	return claims, nil
}
