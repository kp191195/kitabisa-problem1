package main

import "testing"

func TestDoCalc(t *testing.T) {
	var funcResult int64
	var result1 int64
	var result2 int64
	var result3 int64
	str1 := "5 + 2 - 1"
	result1 = 6
	str2 := "100 - 60 + 50 + 50"
	result2 = 140
	str3 := "1023 - 1022 - 45 - 90 + 100"
	result3 = -34

	funcResult = 0
	funcResult, _ = doCalc(str1)
	if funcResult != result1 {
		t.Errorf("doCalc(\"5 + 2 - 1\") failed, expected %d, got %d", result1, funcResult)
	} else {
		t.Logf("doCalc(\"5 + 2 - 1\") success, expected : %d, got : %d", result1, funcResult)
	}

	funcResult, _ = doCalc(str2)
	if funcResult != result2 {
		t.Errorf("doCalc(\"100 - 60 + 50 + 50\") failed, expected %d, got %d", result2, funcResult)
	} else {
		t.Logf("doCalc(\"100 - 60 + 50 + 50\") success, expected : %d, got : %d", result2, funcResult)
	}

	funcResult, _ = doCalc(str3)
	if funcResult != result3 {
		t.Errorf("doCalc(\"1023 - 1022 - 45 - 90 + 100\") failed, expected %d, got %d", result3, funcResult)
	} else {
		t.Logf("doCalc(\"1023 - 1022 - 45 - 90 + 100\") success, expected : %d, got : %d", result3, funcResult)
	}

}
