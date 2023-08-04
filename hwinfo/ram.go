package hwinfo

import (
	"os/exec"
	"strconv"
	"strings"
)

func GetRAM() (float64, error) {
	freeOutput, _ := exec.Command("sh", "-c", "free | grep -w 'Mem:'").Output()
	memInfo := strings.Fields(string(freeOutput))

	memKB, err := strconv.ParseInt(memInfo[1], 10, 64)
	if err != nil {
		return 0, err
	}

	memGB := float64(memKB) / 1024 / 1024

	return memGB, nil
}

func GetSwap() (float64, error) {
	freeOutput, _ := exec.Command("sh", "-c", "free | grep -w 'Swap:'").Output()
	memInfo := strings.Fields(string(freeOutput))

	memKB, err := strconv.ParseInt(memInfo[1], 10, 64)
	if err != nil {
		return 0, err
	}

	memGB := float64(memKB) / 1024 / 1024

	return memGB, nil
}
