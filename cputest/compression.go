package cputest

import (
	"bytes"
	"compress/gzip"
	"crypto/rand"
	"fmt"
	"io"
	"time"
)

func RunCompressionBenchmark() {
	const bufferSize = 500 * 1024 * 1024
	buffer := make([]byte, bufferSize)

	_, err := rand.Read(buffer)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Printf("CPU: gzip compressing 500MB\n")
	startTime := time.Now()

	var compressedData bytes.Buffer
	writer := gzip.NewWriter(&compressedData)

	_, err = io.Copy(writer, bytes.NewReader(buffer))
	if err != nil {
		fmt.Println("Error while compressing:", err)
		return
	}

	err = writer.Close()
	if err != nil {
		return
	}

	elapsedTime := time.Since(startTime)

	fmt.Printf("\t%.3fs\n", elapsedTime.Seconds())
}
