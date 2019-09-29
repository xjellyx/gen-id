package generator

import "testing"

func newGenerator() (ret *GeneratorData) {
	return new(GeneratorData)
}

func TestGeneratorData_GeneratorAddress(t *testing.T) {
	g := newGenerator()
	t.Log(g.GeneratorAddress())
}

func BenchmarkGeneratorData_GeneratorAddress(b *testing.B) {
	for i := 0; i < b.N; i++ {
		g := newGenerator()
		b.Log(g.GeneratorAddress())
	}
}

func TestGeneratorData_GeneratorBankID(t *testing.T) {
	g := new(GeneratorData)
	t.Log(g.GeneratorBankID())
}

func BenchmarkGeneratorData_GeneratorBankID(b *testing.B) {
	for i := 0; i < b.N; i++ {
		g := newGenerator()
		b.Log(g.GeneratorBankID())
	}
}

func TestGeneratorData_GeneratorEmail(t *testing.T) {
	g := new(GeneratorData)
	t.Log(g.GeneratorEmail())
}

func BenchmarkGeneratorData_GeneratorEmail(b *testing.B) {
	for i := 0; i < b.N; i++ {
		g := newGenerator()
		b.Log(g.GeneratorEmail())
	}
}

func TestGeneratorData_GeneratorName(t *testing.T) {
	g := new(GeneratorData)
	t.Log(g.GeneratorName())
}

func BenchmarkGeneratorData_GeneratorName(b *testing.B) {
	for i := 0; i < b.N; i++ {
		g := newGenerator()
		b.Log(g.GeneratorName())
	}
}

func TestGeneratorData_GeneratorPhone(t *testing.T) {
	g := new(GeneratorData)
	t.Log(g.GeneratorPhone())
}

func BenchmarkGeneratorData_GeneratorPhone(b *testing.B) {
	for i := 0; i < b.N; i++ {
		g := newGenerator()
		b.Log(g.GeneratorPhone())
	}
}

func TestGeneratorData_GeneratorProvinceAdnCityRand(t *testing.T) {
	g := new(GeneratorData)
	t.Log(g.GeneratorProvinceAdnCityRand())
}

func BenchmarkGeneratorData_GeneratorProvinceAdnCityRand(b *testing.B) {
	for i := 0; i < b.N; i++ {
		g := newGenerator()
		b.Log(g.GeneratorProvinceAdnCityRand())
	}
}

func TestGeneratorData_GeneratorIDCart(t *testing.T) {
	g := new(GeneratorData)
	t.Log(g.GeneratorIDCart(nil))
}

func BenchmarkGeneratorData_GeneratorIDCart(b *testing.B) {
	for i := 0; i < b.N; i++ {
		g := newGenerator()
		b.Log(g.GeneratorIDCart(nil))
	}
}
