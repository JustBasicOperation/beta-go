package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
)

// generateRSAKey 生成RSA公私钥
func generateRSAKey() ([]byte, []byte) {
	priKey, _ := rsa.GenerateKey(rand.Reader, 1024)
	derStream := x509.MarshalPKCS1PrivateKey(priKey)
	priBlock := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: derStream,
	}
	privateKey := pem.EncodeToMemory(priBlock)
	pubKey := &priKey.PublicKey

	derPkix, _ := x509.MarshalPKIXPublicKey(pubKey)
	pubBlock := &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: derPkix,
	}
	publicKey := pem.EncodeToMemory(pubBlock)
	fmt.Println(string(privateKey))
	fmt.Println(string(publicKey))
	return privateKey, publicKey
}

// RSAEncrypt rsa 加密
func RSAEncrypt(publicKey, plainText []byte) ([]byte, error) {
	block, _ := pem.Decode(publicKey) // 将公钥解析成公钥实例
	pubInst, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return rsa.EncryptPKCS1v15(rand.Reader, pubInst.(*rsa.PublicKey), plainText)
}

// RSADeCrypt rsa 解密
func RSADeCrypt(privateKey, cipherText []byte) ([]byte, error) {
	block, _ := pem.Decode(privateKey) // 将私钥解析成私钥实例
	if block == nil {
		return nil, errors.New("private key error")
	}
	priInst, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return rsa.DecryptPKCS1v15(rand.Reader, priInst, cipherText)
}

// AESEncrypt aes 加密
func AESEncrypt(key, plainText []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	iv := make([]byte, block.BlockSize())
	ctrEncryptor := cipher.NewCTR(block, iv)
	encrypted := make([]byte, len(plainText))
	ctrEncryptor.XORKeyStream(encrypted, plainText)
	return encrypted, nil
}

// AESDecrypt aes 解密
func AESDecrypt(key, cipherText []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	iv := make([]byte, block.BlockSize())
	ctrDecryptor := cipher.NewCTR(block, iv)
	decrypted := make([]byte, len(cipherText))
	ctrDecryptor.XORKeyStream(decrypted, cipherText)
	return decrypted, nil
}
