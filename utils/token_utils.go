package utils

import (
	"crypto/rsa"
	"encoding/pem"
	"fmt"

	"github.com/dgrijalva/jwt-go"
)

type CustomClaims struct {
	Sid string `json:"sid"`
	jwt.StandardClaims
}

func parseRSAPublicKeyFromPEM(publicKeyPEM string) (*rsa.PublicKey, error) {
	block, _ := pem.Decode([]byte(publicKeyPEM))
	if block == nil {
		return nil, fmt.Errorf("failed to parse PEM block containing the public key")
	}
	pub, err := jwt.ParseRSAPublicKeyFromPEM([]byte(publicKeyPEM))
	if err != nil {
		return nil, fmt.Errorf("failed to parse public key: %v", err)
	}
	return pub, nil
}

func GetSessionIDFromToken(tokenString string) (string, error) {
	claims := &CustomClaims{}

	publicKeyPEM := `-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA2nwJYO3s0w62+pK+Yuz0
p3S/CZB4kNssmzduFWNmASBID3+iLU+ELqw3BAjKCbAfKdRjxW/L5vCc2qQDilQE
UqB+LCFbiCahyk1tqRBgfAMe2yu+h33XZceAZAQJNic0uPMJ67Ljm7Iy19opPvaU
1hbVwu9MBYGxdRoycoA7o0vDYMme3foG2DO0o8aH44Th0YbUdFBsOwDZDvaPAwSI
gc0aydG28Y0n2aOSvpLWtEWyFn1iOHdI0Z+Y0oTx3HlwpCj/ele2k07U50jR9TOV
mQzte20LdNbDNK/OSCRfsUTtNQSegmLDcpFmKDUr2JjdxiP79zQ9QxsaprYxiLS9
3QIDAQAB
-----END PUBLIC KEY-----`

	publicKey, err := parseRSAPublicKeyFromPEM(publicKeyPEM)
	if err != nil {
		return "", fmt.Errorf("error parsing RSA public key: %v", err)
	}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return publicKey, nil
	})

	if err != nil {
		return "", fmt.Errorf("error decoding JWT: %v", err)
	}

	if token.Valid {
		return claims.Sid, nil
	}

	return "", fmt.Errorf("invalid token")

}
