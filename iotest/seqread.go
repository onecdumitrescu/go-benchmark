package iotest

import (
	"fmt"
	"io"
	"os"
	"time"
)

const (
	fileSizeInMiB = 1024
	filePath      = "testfile.bin"
)

func createTempFile() *os.File {
	buffer := make([]byte, fileSizeInMiB*1024*1024)

	file, err := os.Create(filePath)
	if err != nil {
		fmt.Println("Error creating the temp file:", err)
		return nil
	}

	_, err = file.Write(buffer)
	if err != nil {
		fmt.Println("Error writing the block to disk:", err)
		return nil
	}

	err = file.Sync()
	if err != nil {
		fmt.Println("Error syncing the file to disk:", err)
		return nil
	}

	err = file.Close()
	if err != nil {
		return nil
	}

	return file
}

func readTempFile() {
	fmt.Printf("\nIO: Sequential read\n")

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening the file:", err)
		return
	}

	defer closeAndRemoveFile(file)

	buffer := make([]byte, 64*1024)

	startTime := time.Now()
	bytesRead := int64(0)

	for {
		n, err := file.Read(buffer)
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("Error reading from the file:", err)
			return
		}
		bytesRead += int64(n)
	}

	elapsedTime := time.Since(startTime)

	readSpeed := float64(bytesRead) / (1024 * 1024) / elapsedTime.Seconds()

	fmt.Printf("\tRead speed:\t%.2f MiB/s\n", readSpeed)
}

func closeAndRemoveFile(file *os.File) {
	err := file.Close()
	if err != nil {
		fmt.Println("Error closing the temp file:", err)
		return
	}

	err = os.Remove(filePath)
	if err != nil {
		fmt.Println("Error deleting the temp file:", err)
		return
	}
}

func RunSequentialReadBenchmark() {
	if file := createTempFile(); file != nil {
		readTempFile()
	}
}
