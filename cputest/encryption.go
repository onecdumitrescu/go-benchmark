package cputest

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"fmt"
	"time"
)

func RunEncryptionBenchmark() {
	data := make([]byte, 500*1024*1024)
	key := bytes.Repeat([]byte{0x01}, 32)
	block, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println("Error while encrypting:", err)
		return
	}

	ciphertext := make([]byte, aes.BlockSize+len(data))
	iv := ciphertext[:aes.BlockSize]

	mode := cipher.NewCBCEncrypter(block, iv)

	fmt.Printf("CPU: AES encrypting 500MB\n")
	startTime := time.Now()

	mode.CryptBlocks(ciphertext[aes.BlockSize:], data)
	elapsedTime := time.Since(startTime)

	fmt.Printf("\t%.3fs\n", elapsedTime.Seconds())
}
