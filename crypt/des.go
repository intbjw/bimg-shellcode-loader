package crypt

import (
	"bytes"
	"crypto/cipher"
	"crypto/des"
)

// 使用DES CFB模式进行加密
func DesEncrypt(origData, key, iv []byte) ([]byte, error) {
	// 创建一个DES密码分组
	block, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}

	// 对原始数据进行填充
	origData = pkcs7Padding(origData, block.BlockSize())

	// 创建一个CFB加密模式
	cfb := cipher.NewCFBEncrypter(block, iv)

	// 加密数据
	encrypted := make([]byte, len(origData))
	cfb.XORKeyStream(encrypted, origData)

	return encrypted, nil
}

// 使用DES CFB模式进行解密
func DesDecrypt(encryptedData, key, iv []byte) ([]byte, error) {
	// 创建一个DES密码分组
	block, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}

	// 创建一个CFB解密模式
	cfb := cipher.NewCFBDecrypter(block, iv)

	// 解密数据
	decrypted := make([]byte, len(encryptedData))
	cfb.XORKeyStream(decrypted, encryptedData)

	// 对解密后的数据进行去填充
	decrypted = pkcs7UnPadding(decrypted)

	return decrypted, nil
}

// PKCS7填充
func pkcs7Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

// PKCS7去填充
func pkcs7UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}
