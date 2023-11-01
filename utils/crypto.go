package utils

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"log"
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

// pkcs5Padding 填充
// 当明文长度不够时，缺几位填几个几
func pkcs5Padding(data []byte, blockSize int) []byte {
	//判断缺少几位长度。最少1，最多 blockSize
	padding := blockSize - len(data)%blockSize
	//补足位数。把切片[]byte{byte(padding)}复制padding个
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, padText...)
}

// pkcs5UnPadding 填充的反向操作
func pkcs5UnPadding(data []byte) ([]byte, error) {
	length := len(data)
	if length == 0 {
		return nil, errors.New("加密字符串错误！")
	}
	//获取填充的个数
	unPadding := int(data[length-1])
	return data[:(length - unPadding)], nil
}

// AesEncryptCBC 加密
func AesEncryptCBC(iv, data []byte, key []byte) ([]byte, error) {
	//创建加密实例
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	//判断加密快的大小
	blockSize := block.BlockSize()
	//填充
	encryptBytes := pkcs5Padding(data, blockSize)
	log.Printf("encryptBytes: %v", encryptBytes)
	//初始化加密数据接收切片
	crypted := make([]byte, len(encryptBytes))
	//使用cbc加密模式
	blockMode := cipher.NewCBCEncrypter(block, iv)
	//执行加密
	blockMode.CryptBlocks(crypted, encryptBytes)
	return crypted, nil
}

// AesDecryptCBC 解密
func AesDecryptCBC(iv, data []byte, key []byte) ([]byte, error) {
	//创建实例
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	//使用cbc
	blockMode := cipher.NewCBCDecrypter(block, iv)
	//初始化解密数据接收切片
	crypted := make([]byte, len(data))
	//执行解密
	blockMode.CryptBlocks(crypted, data)
	//去除填充
	crypted, err = pkcs5UnPadding(crypted)
	if err != nil {
		return nil, err
	}
	return crypted, nil
}
