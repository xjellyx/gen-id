package gen_id

import "testing"

func TestNewGeneratorData(t *testing.T) {
	t.Log(NewGeneratorData())
}

func BenchmarkNewGeneratorData(b *testing.B) {
	for i:=0;i<b.N;i++{
		b.Log(NewGeneratorData())
	}
}
