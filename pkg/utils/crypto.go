package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"

	"github.com/DoWithLogic/go-echo-realworld/config"
)

func Encrypt(pwd string, cfg config.Config) (string, error) {
	return encrypt(pwd, []byte(cfg.Authentication.SecretKey), []byte(cfg.Authentication.SaltKey))
}

func Decrypt(pwd string, cfg config.Config) (string, error) {
	return decrypt(pwd, []byte(cfg.Authentication.SecretKey), []byte(cfg.Authentication.SaltKey))
}

func encrypt(text string, key, salt []byte) (string, error) {
	plaintext := []byte(text)

	// Create a new AES cipher block
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	// Create a GCM (Galois/Counter Mode) cipher using AES
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	// Create a nonce by concatenating salt and random bytes. Nonce must be unique for each encryption
	nonce := make([]byte, gcm.NonceSize())
	copy(nonce, salt)

	// Encrypt the data using AES-GCM
	ciphertext := gcm.Seal(nil, nonce, plaintext, nil)

	// Include the nonce in the encrypted data
	encryptedData := append(nonce, ciphertext...)

	return base64.StdEncoding.EncodeToString(encryptedData), nil
}

func decrypt(encryptedText string, key, salt []byte) (string, error) {
	// Decode base64
	encryptedData, err := base64.StdEncoding.DecodeString(encryptedText)
	if err != nil {
		return "", err
	}

	// Create a new AES cipher block
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	// Create a GCM (Galois/Counter Mode) cipher using AES
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	// Nonce size is determined by the choice of GCM mode and its associated size for the given key
	nonceSize := gcm.NonceSize()

	// Extract the nonce from the encrypted data
	nonce, encryptedMessage := encryptedData[:nonceSize], encryptedData[nonceSize:]

	// Decrypt the data using AES-GCM
	plaintext, err := gcm.Open(nil, nonce, encryptedMessage, nil)
	if err != nil {
		return "", err
	}

	return string(plaintext), nil
}
