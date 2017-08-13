package main

import (
	"fmt"
	"strconv"
)

const annotation = "black-death"

func getBlackDeath(annotations map[string]string) (int64, error) {
	if val, ok := annotations[annotation]; ok {
		return strconv.ParseInt(val, 10, 64)
	}

	return 0, fmt.Errorf("annotation does not exist")
}
