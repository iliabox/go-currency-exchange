package vo

import (
	"fmt"
	"strconv"
	"strings"
)

type Number string

func NewNumber(v string) (Number, error) {
	if !isNumeric(v) {
		return "", fmt.Errorf("invalid value")
	}
	f, err := strconv.ParseFloat(v, 64)
	if err != nil || f <= 0 {
		return "", fmt.Errorf("invalid value")
	}
	if len(v) >= 32 {
		return "", fmt.Errorf("invalid value")
	}
	return Number(v), nil
}

func (n Number) String() string { return string(n) }

func isNumeric(v string) bool {
	v = strings.TrimSpace(v)
	_, err := strconv.ParseFloat(v, 64)
	return err == nil
}
