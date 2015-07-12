package main

import "testing"

func TestFoo(t *testing.T) {
	t.Log("Running")
	if false == true {
		t.Error("The world is a strange place")
	}
}
