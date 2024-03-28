package utils

import (
	"fmt"
	"strings"
)

func ConvertVectorToString(vector []float64) string {
	var sb strings.Builder
	sb.WriteRune('[')
	for i, v := range vector {
		if i > 0 {
			sb.WriteString(", ")
		}
		sb.WriteString(fmt.Sprintf("%f", v))
	}
	sb.WriteRune(']')
	return sb.String()
}
