package main

import (
    "crypto/aes"
    "crypto/cipher"
    "crypto/rand"
    "crypto/sha256"
    "encoding/base64"
    "errors"
    "io"

    "golang.org/x/crypto/bcrypt"
    "golang.org/x/crypto/pbkdf2"
)

func HashMasterPassword(password string) (string, error) {
    hash, err := bcrypt.GenerateFromPassword([]byte(password), 12)
    return string(hash), err
}

func VerifyMasterPassword(hash string, password string) error {
    return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}

func DeriveKey(password string, salt []byte) []byte {
    return pbkdf2.Key([]byte(password), salt, 150000, 32, sha256.New)
}

func EncryptAES(plaintext []byte, key []byte) (ciphertext, nonce []byte, err error) {
    block, err := aes.NewCipher(key)
    if err != nil {
        return nil, nil, err
    }

    gcm, err := cipher.NewGCM(block)
    if err != nil {
        return nil, nil, err
    }

    nonce = make([]byte, gcm.NonceSize())
    if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
        return nil, nil, err
    }

    ciphertext = gcm.Seal(nil, nonce, plaintext, nil)
    return ciphertext, nonce, nil
}

func DecryptAES(ciphertext, nonce, key []byte) ([]byte, error) {
    block, err := aes.NewCipher(key)
    if err != nil {
        return nil, err
    }

    gcm, err := cipher.NewGCM(block)
    if err != nil {
        return nil, err
    }

    if len(nonce) != gcm.NonceSize() {
        return nil, errors.New("invalid nonce size")
    }

    plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
    return plaintext, err
}

func EncodeBase64(b []byte) string {
    return base64.StdEncoding.EncodeToString(b)
}

func DecodeBase64(s string) ([]byte, error) {
    return base64.StdEncoding.DecodeString(s)
}
