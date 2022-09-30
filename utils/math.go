package utils

import (
	"bytes"
	"encoding/gob"
	"math"
	"strconv"
)

const float64EqualityThreshold = 1e-9

func AlmostEqual(a, b float64) bool {
	return math.Abs(a-b) <= float64EqualityThreshold
}

func ParseFloat64(s string) float64 {
	f, _ := strconv.ParseFloat(s, 64)
	return f
}

func ParseInt(s string) int {
	i, _ := strconv.ParseInt(s, 10, 64)
	return int(i)
}

func ParseInt64(s string) int64 {
	i, _ := strconv.ParseInt(s, 10, 64)
	return i
}

func ParseBool(s string) bool {
	b, _ := strconv.ParseBool(s)
	return b
}

func Hash(s interface{}) []byte {
	var b bytes.Buffer
	gob.NewEncoder(&b).Encode(s)
	return b.Bytes()
}
