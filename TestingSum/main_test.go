package main

import "testing"

func Test_CreateSum(t *testing.T) {
	result := CreateSum(2, 3)
	if result != 5 {
		t.Error("incorrect result: expected 5, got", result)
	}
}
