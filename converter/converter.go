package converter

import (
	"log"
	"strconv"
	"strings"
)

// ConvertToFloat64 converts a string to float64. If it fails, zero is returned.
func ConvertToFloat64(value string) float64 {
	result, err := strconv.ParseFloat(strings.TrimSpace(value), 64)

	if err != nil {
		log.Printf("Error %v. Returning value = 0", err)
		return 0
	}

	return result
}
