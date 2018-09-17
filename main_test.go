package main

import (
	"testing"
)

func getTestData() map[string]int {
	testData := map[string]int{
		"":             -1,
		")":            0,
		"(":            0,
		"))":           0,
		"(1234":        0,
		"([}]){}":      2,
		"()()":         -1,
		"({[]}{})":     -1,
		"(0[1]{2})":    -1,
		"([](){([])})": -1,
		"()[]}":        4,
		"{{[()]]":      6,
		"{{{[][][]":    2,
		"{*{{}":        2,
		"[[*":          1,
		"{*}":          -1,
		"{{":           1,
		"{}":           -1,
		"}":            0,
		"*{}":          -1,
		"{{{**[][][]":  2,
	}

	return testData
}

func TestFindWrongBrackets(t *testing.T) {
	testData := getTestData()

	for dataKey, dataValue := range testData {
		value := findWrongBrackets(dataKey)

		if value != dataValue {
			t.Errorf("Test failed - results not match\nfor key '%v'\ngot: %v\nexpected: %v", dataKey, value, dataValue)
		}
	}
}
