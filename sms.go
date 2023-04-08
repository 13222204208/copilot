package copilot

import (
	"fmt"
	"math/rand"
)

// GenerateCode generates a random 6 digit code
func GenerateCode() string {
	generator := rand.New(rand.NewSource(1000))
	randomNum := generator.Intn(900000) + 100000
	return fmt.Sprintf("%d", randomNum)
}
