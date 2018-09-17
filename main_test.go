package main

import (
	"testing"
)

func getTestData() map[string]int {
	testData := map[string]int{
		// wrong brackets
		")":       0,
		"(":       0,
		"))":      0,
		"(1234":   0,
		"([}]){}": 2,

		// right brackets
		"":          -1,
		"()()":      -1,
		"({[]}{})":  -1,
		"(0[1]{2})": -1,
	}

	return testData
}

func TestFindWrongBrackets(t *testing.T) {
	testData := getTestData()

	for dataKey, dataValue := range testData {
		value := findWrongBrackets(dataKey)

		if value != dataValue {
			t.Error("Test failed - wrong value for key:" + dataKey)
		}
	}
}

func TestCheckBrackets(t *testing.T) {

}
