package cputest

import (
	"bytes"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"io"
	"time"
)

func RunHashingBenchmark() {
	const bufferSize = 500 * 1024 * 1024
	buffer := make([]byte, bufferSize)

	_, err := rand.Read(buffer)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Printf("CPU: SHA-256 hashing 500MB\n")
	startTime := time.Now()

	h := sha256.New()
	_, err = io.Copy(h, bytes.NewReader(buffer))
	if err != nil {
		fmt.Println("Error while hashing:", err)
		return
	}

	_ = h.Sum(nil)

	elapsedTime := time.Since(startTime)

	fmt.Printf("\t%.3fs\n", elapsedTime.Seconds())
}
