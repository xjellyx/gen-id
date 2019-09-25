package gen_id

import "github.com/srlemon/gen-id/generator"

// NewGeneratorData
func NewGeneratorData()(ret *generator.GeneratorData)  {
	var(
		data = new(GeneratorData)
	)
	data.GeneratorBankID()
	data.GeneratorAddress()
	data.GeneratorEmail()
	data.GeneratorIDCart()
	data.GeneratorName()
	data.GeneratorPhone()
	ret  =data
	return
}
