package iotest

import (
	"fmt"
	"os"
	"time"
)

func RunSequentialWriteBenchmark() {
	n := 3
	fileSizeInMiB := 1024
	filePath := "testfile.bin"

	buffer := make([]byte, 64*1024)

	totalSpeed := float64(0)

	fmt.Printf("\nIO: Sequential write\n")

	for i := 1; i <= n; i++ {
		file, err := os.Create(filePath)
		if err != nil {
			fmt.Println("Error creating the temp file:", err)
			return
		}

		startTime := time.Now()
		bytesWritten := int64(0)

		for b := 0; b < fileSizeInMiB; b++ {
			buf, err := file.Write(buffer)
			if err != nil {
				fmt.Println("Error writing to the temp file:", err)
				return
			}
			bytesWritten += int64(buf)

			err = file.Sync()
			if err != nil {
				fmt.Println("Error writing the block to disk:", err)
				return
			}
		}

		elapsedTime := time.Since(startTime)

		err = file.Close()
		if err != nil {
			return
		}

		err = os.Remove(filePath)
		if err != nil {
			fmt.Println("Error deleting the temp file:", err)
			return
		}

		writeSpeed := float64(bytesWritten) / (1024 * 1024) / elapsedTime.Seconds()
		totalSpeed += writeSpeed

		fmt.Printf("\tIteration %d write speed: %.2f MiB/s\n", i, writeSpeed)
	}

	fmt.Printf("\tAverage write speed:\t%.2f MiB/s\n", totalSpeed/float64(n))
}
