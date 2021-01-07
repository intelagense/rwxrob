package lib

import "testing"

func BenchmarkNewPerson(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewPerson()
	}
}

func BenchmarkNewPersonI(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewPersonI()
	}
}

func BenchmarkNewPersonII(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewPersonII()
	}
}
