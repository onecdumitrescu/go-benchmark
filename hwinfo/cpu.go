package hwinfo

import (
	"log"
	"os"
	"runtime"
	"strconv"
	"strings"
)

func GetCPUName() string {
	bytes, err := os.ReadFile("/proc/cpuinfo")
	if err != nil {
		log.Fatal(err)
	}

	cpuinfo := string(bytes)
	startIdx := strings.Index(cpuinfo, "model name") + 13
	endIdx := strings.Index(cpuinfo[startIdx:], "\n") + startIdx

	return cpuinfo[startIdx:endIdx]
}

func GetCPUSpeed() string {
	bytes, err := os.ReadFile("/proc/cpuinfo")
	if err != nil {
		log.Fatal(err)
	}

	cpuinfo := string(bytes)
	startIdx := strings.Index(cpuinfo, "cpu MHz") + 11
	endIdx := strings.Index(cpuinfo[startIdx:], "\n") + startIdx

	return cpuinfo[startIdx:endIdx]
}

func GetCPUCores() string {
	return strconv.Itoa(runtime.NumCPU())
}
