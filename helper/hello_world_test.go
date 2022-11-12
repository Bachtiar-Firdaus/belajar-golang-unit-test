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

// example benchmark
func BenchmarkHelloWorldV1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWord("bachtiar")
	}
}

func BenchmarkHelloWorldV2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWord("firdaus")
	}
}
