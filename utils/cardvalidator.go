package utils

import (
	"strconv"
)

type CardValidator struct {
}

func NewCardValidator() *CardValidator {
	return &CardValidator{}
}

func (cv *CardValidator) FirstNumber(value string) string {
	var val string
	for i, char := range value {
		if i == 0 {
			val = string(byte(char))
		}
	}
	return val
}
func (cv *CardValidator) InRange(value, lower, upper string) bool {
	intValue, _ := strconv.Atoi(value)
	intLower, _ := strconv.Atoi(lower)
	intUpper, _ := strconv.Atoi(upper)
	return intLower < intValue && intValue > intUpper
}
