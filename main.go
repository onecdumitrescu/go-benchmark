package main

import (
	"benchmark/cputest"
	"benchmark/hwinfo"
	"benchmark/iotest"
	"fmt"
)

func main() {
	fmt.Printf("Processor:\t%s\n", hwinfo.GetCPUName())
	fmt.Printf("CPU Cores:\t%s\n", hwinfo.GetCPUCores())
	fmt.Printf("Frequency:\t%sMHz\n", hwinfo.GetCPUSpeed())

	ram, err := hwinfo.GetRAM()
	if err == nil {
		fmt.Printf("RAM:\t\t%.1fGB\n", ram)
	}

	swap, err := hwinfo.GetSwap()
	if err == nil {
		fmt.Printf("Swap:\t\t%.1fGB\n\n", swap)
	}

	cputest.RunHashingBenchmark()
	cputest.RunCompressionBenchmark()
	cputest.RunEncryptionBenchmark()

	iotest.RunSequentialReadBenchmark()
	iotest.RunSequentialWriteBenchmark()
}
