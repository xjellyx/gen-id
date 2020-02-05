package genid

import "github.com/olongfen/gen-id/generator"

// NewGeneratorData
func NewGeneratorData(isFullAge *bool) (ret *generator.GeneratorData) {
	var (
		data = new(generator.GeneratorData)
	)
	data.GeneratorBankID()
	data.GeneratorAddress()
	data.GeneratorEmail()
	data.GeneratorIDCart(isFullAge)
	data.GeneratorName()
	data.GeneratorPhone()
	ret = data
	return
}
