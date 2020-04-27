package datasources

import (
	"crypto/ecdsa"
	"crypto/rsa"
	"io/ioutil"

	jwt "github.com/dgrijalva/jwt-go"
)

var (
	// JWTECVerifyKey JWT public ECDSA key
	JWTECVerifyKey *ecdsa.PublicKey
	// JWTECSignKey JWT private ECDSA key
	JWTECSignKey *ecdsa.PrivateKey

	// JWTRSAVerifyKey JWT public RSA key
	JWTRSAVerifyKey *rsa.PublicKey
	// JWTRSASignKey JWT private RSA key
	JWTRSASignKey *rsa.PrivateKey

	// JWTHSSignKey JWT private HS key
	JWTHSSignKey *[]byte
)

// ValidationJWT JWT validation func
var ValidationJWT jwt.Keyfunc = func(token *jwt.Token) (interface{}, error) {
	return JWTECVerifyKey, nil
}

// NewJWTKey new JWT key
func NewJWTKey(jwtSecret string) (err error) {
	*JWTHSSignKey = []byte(jwtSecret)
	return
}

// NewJWTRSAKey new JWT key
func NewJWTRSAKey(pubKeyPath string, privKeyPath string) (err error) {
	signByte, err := ioutil.ReadFile(privKeyPath)
	if err != nil {
		return
	}

	JWTRSASignKey, err = jwt.ParseRSAPrivateKeyFromPEM(signByte)
	if err != nil {
		return
	}

	verifyByte, err := ioutil.ReadFile(pubKeyPath)
	if err != nil {
		return
	}

	JWTRSAVerifyKey, err = jwt.ParseRSAPublicKeyFromPEM(verifyByte)
	if err != nil {
		return
	}

	return
}

// NewJWTECKey new JWT key
func NewJWTECKey(pubKeyPath string, privKeyPath string) (err error) {
	signByte, err := ioutil.ReadFile(privKeyPath)
	if err != nil {
		return
	}

	JWTECSignKey, err = jwt.ParseECPrivateKeyFromPEM(signByte)
	if err != nil {
		return
	}

	verifyByte, err := ioutil.ReadFile(pubKeyPath)
	if err != nil {
		return
	}

	JWTECVerifyKey, err = jwt.ParseECPublicKeyFromPEM(verifyByte)
	if err != nil {
		return
	}

	return
}
