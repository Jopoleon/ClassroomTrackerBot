package github_api

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"os"
	"time"
)

// loadPrivateKey loads the RSA private key from a PEM file
func loadPrivateKey(path string) (*rsa.PrivateKey, error) {
	privateKeyData, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	pemBlock, _ := pem.Decode(privateKeyData)
	if pemBlock == nil {
		return nil, fmt.Errorf("could not parse PEM block")
	}

	return x509.ParsePKCS1PrivateKey(pemBlock.Bytes)
}

// createJWT creates a JWT for GitHub App authentication
func createJWT(appID int64, privateKey *rsa.PrivateKey) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"iat": time.Now().Unix(),
		"exp": time.Now().Add(time.Minute * 10).Unix(),
		"iss": appID,
	})

	return token.SignedString(privateKey)
}
func CreateJWTToken(privateKeyFilePath string) (string, error) {

	// Load the private key from your file
	privateKey, err := loadPrivateKey(privateKeyFilePath)

	if err != nil {
		fmt.Println("Error loading private key:", err)
		return "", err
	}

	// Create a JWT token
	jwtToken, err := createJWT(696840, privateKey) // Replace with your App ID
	if err != nil {
		fmt.Println("Error creating JWT:", err)
		return "", err
	}
	return jwtToken, err

}
