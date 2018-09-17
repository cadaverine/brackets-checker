package main

import (
	"testing"
)

func getTestData() map[string]int {
	var testData map[string]int

	// right brackets
	testData["()()"] = -1
	testData[""] = -1
	testData["({[]}{})"] = -1

	// wrong brackets
	testData[")"] = 0
	testData["))"] = 0
	testData["("] = 0
	testData[")"] = 0
	testData["([}]){}"] = 2

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
