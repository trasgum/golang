package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

//var test_cases map[string]int{
//	"": 0,
//	"1": 1,
//	"1,2": 3,
//	"1,2,10,100": 113,
//	"1\n2,3": 6,
//	"//;\n100;10": 110,
//}

func TestShouldReturnZeroForEmptyString(t *testing.T) {
	suma, _ := Add("")
	assert.Equal(t, int64(0),
		suma)
}

func TestShouldReturnNumberForOneNumberString(t *testing.T) {
	suma, _ := Add("1")
	assert.Equal(t, int64(1),
		suma)
}

func TestShouldReturnSumNumbersForTwoNumberString(t *testing.T) {
	suma, _ := Add("1,2")
	assert.Equal(t, int64(3),
		suma)
}

func TestShouldReturnSumNumbersForAnyNumberString(t *testing.T) {
	suma, _ := Add("1,2,10,100")
	assert.Equal(t, int64(113),
		suma)
}

func TestShouldReturnSumNumbersForNewLineNumberString(t *testing.T) {
	suma, _ := Add("1\n2,3")
	assert.Equal(t, int64(6),
		suma)
}

func TestShouldReturnSumNumbersFoCustomDelimiterString(t *testing.T) {
	suma, _ := Add("//;\n100;10")
	assert.Equal(t, int64(110),
		suma)
}

func TestShouldReturnErrorForNegativeNumber(t *testing.T) {
	assert.Panics(t, func() { Add("-1")})
}