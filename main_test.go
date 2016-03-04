package main

import (
	"testing"
)

func TestSample2(t *testing.T) {
	t.Error("test")
}

func TestSample(t *testing.T) {
	if false {
		t.Error("sample test error")
	}
}
