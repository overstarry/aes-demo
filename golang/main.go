package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"fmt"
)

func AesEncrypt(encryptStr string, key []byte, iv string) (string, error) {
	encryptBytes := []byte(encryptStr)
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	blockSize := block.BlockSize()
	encryptBytes = pkcs5Padding(encryptBytes, blockSize)

	blockMode := cipher.NewCBCEncrypter(block, []byte(iv)[:block.BlockSize()])
	encrypted := make([]byte, len(encryptBytes))
	blockMode.CryptBlocks(encrypted, encryptBytes)
	return hex.EncodeToString(encrypted), nil
}

func AesDecrypt(decryptStr string, key []byte, iv string) (string, error) {
	decryptBytes, err := hex.DecodeString(decryptStr)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	blockMode := cipher.NewCBCDecrypter(block, []byte(iv)[:block.BlockSize()])
	decrypted := make([]byte, len(decryptBytes))

	blockMode.CryptBlocks(decrypted, decryptBytes)
	decrypted = pkcs5UnPadding(decrypted)
	return string(decrypted), nil
}

func pkcs5Padding(cipherText []byte, blockSize int) []byte {
	padding := blockSize - len(cipherText)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(cipherText, padText...)
}

func pkcs5UnPadding(decrypted []byte) []byte {
	length := len(decrypted)
	unPadding := int(decrypted[length-1])
	return decrypted[:(length - unPadding)]
}
func main() {
	data := "i am test data"
	key := []byte("abcdefghabcdefgh")
	iv := "1234567812345678"
	encrypt, err := AesEncrypt(data, key, iv)
	if err != nil {
		panic(err)
		return
	}
	fmt.Println(encrypt)
	decrypt, err := AesDecrypt(encrypt, key, iv)
	if err != nil {
		panic(err)
	}
	fmt.Println(decrypt)
}
