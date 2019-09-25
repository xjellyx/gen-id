package idSDK

import (
	"fmt"
	"github.com/srlemon/idSDK/generator"
)

func run()  {
	// 生成总的信息
	fmt.Println(generator.NewGeneratorData())
	// 分个单独获取
	g:=new(generator.GeneratorData)
	fmt.Println(g.GeneratorPhone())
	fmt.Println(g.GeneratorName())
	fmt.Println(g.GeneratorIDCart())
	fmt.Println(g.GeneratorEmail())
	fmt.Println(g.GeneratorBankID())
	fmt.Println(g.GeneratorAddress())
}
