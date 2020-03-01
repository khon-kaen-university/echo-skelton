package datasources

import (
	"crypto/ecdsa"
	"io/ioutil"

	jwt "github.com/dgrijalva/jwt-go"
)

var (
	// JWTVerifyKey JWT public key
	JWTVerifyKey *ecdsa.PublicKey
	// JWTSignKey JWT private key
	JWTSignKey *ecdsa.PrivateKey
)

// ValidationJWT JWT validation func
var ValidationJWT jwt.Keyfunc = func(token *jwt.Token) (interface{}, error) {
	return JWTVerifyKey, nil
}

// NewJWTKey new JWT key
func NewJWTKey(pubKeyPath string, privKeyPath string) (err error) {
	signByte, err := ioutil.ReadFile(privKeyPath)
	if err != nil {
		return
	}

	JWTSignKey, err = jwt.ParseECPrivateKeyFromPEM(signByte)
	if err != nil {
		return
	}

	verifyByte, err := ioutil.ReadFile(pubKeyPath)
	if err != nil {
		return
	}

	JWTVerifyKey, err = jwt.ParseECPublicKeyFromPEM(verifyByte)
	if err != nil {
		return
	}

	return
}
