package genid

import "testing"

func TestNewGeneratorData(t *testing.T) {
	t.Log(NewGeneratorData(nil))
}

func BenchmarkNewGeneratorData(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.Log(NewGeneratorData(nil))
	}
}
