package helper

import "testing"

// example 1

func TestHelloWorld(t *testing.T) {
	result := HelloWord("Daus")
	if result != "Hello Daus" {
		panic("Result is not Hello Daus")
	}
}

// example 2
func TestHelloWorldV2(t *testing.T) {
	result := HelloWord("Bachtiar")
	if result != "Hello Bachtiar" {
		panic("Result is not Hello Bachtiar")
	}
}
