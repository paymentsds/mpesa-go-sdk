package mpesa

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
)

// TokenGenerator is an interface for generating and retrieving access tokens.
type TokenGenerator interface {
	GenerateAccessToken() error
	GetAccessToken() string
}

// TokenGeneratorImpl is an implementation of the TokenGenerator interface.
// It generates an access token by encrypting the api key with the public key
// using the RSA algorithm.
type TokenGeneratorImpl struct {
	publicKey   string // the public key
	apiKey      string // the api key
	accessToken string // the generated access token
}

// NewTokenGenerator creates a new TokenGeneratorImpl.
func NewTokenGenerator() *TokenGeneratorImpl {
	return &TokenGeneratorImpl{}
}

// SetApiKey sets the api key for the TokenGeneratorImpl.
func (tok *TokenGeneratorImpl) SetApiKey(apiKey string) {
	tok.apiKey = apiKey
}

// SetPublicKey sets the public key for the TokenGeneratorImpl.
func (tok *TokenGeneratorImpl) SetPublicKey(publicKey string) {
	tok.publicKey = publicKey
}

// GetAccessToken retrieves the generated access token.
func (tok *TokenGeneratorImpl) GetAccessToken() string {

	return tok.accessToken
}

// GenerateAccessToken generates an access token by encrypting the api key
// with the public key using the RSA algorithm.
func (tok *TokenGeneratorImpl) GenerateAccessToken() error {
	if tok.apiKey == "" || tok.publicKey == "" {
		return errors.New("api_key and public_key are required to generate an access token")
	}

	formattedPublicKey, err := tok.formatPublicKey(tok.publicKey)
	if err != nil {
		return err
	}

	block, _ := pem.Decode([]byte(formattedPublicKey))
	if block == nil {
		return errors.New("invalid public key")
	}

	rsaPublicKey, err := x509.ParsePKIXPublicKey(

		block.Bytes)
	if err != nil {
		return err
	}
	encryptedAPIKey, err := rsa.EncryptPKCS1v15(nil, rsaPublicKey.(*rsa.PublicKey), []byte(tok.apiKey))
	if err != nil {
		return err
	}

	tok.accessToken = base64.StdEncoding.EncodeToString(encryptedAPIKey)

	return nil
}
func (conf *TokenGeneratorImpl) formatPublicKey(publicKey string) (string, error) {
	header := "-----BEGIN PUBLIC KEY-----"
	footer := "-----END PUBLIC KEY-----"

	formattedPublicKey := header + "\n" + publicKey + "\n" + footer

	return formattedPublicKey, nil
}
