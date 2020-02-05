# gen-id
一个身份证、名字、邮箱、地址、手机号码等随机生成的sdk

# Installation
`go get github.com/olongfen/gen-id`

如果网速过慢:
```
export GO111MODULE=on
export GOPROXY=https://goproxy.io
go get github.com/olongfen/gen-id
```

# Example

```golang
package main

import (
	"fmt"
	"github.com/olongfen/gen-id"
	"github.com/olongfen/gen-id/generator"
)

func main()  {
	// 生成总的信息
	fmt.Println(gen_id.NewGeneratorData(nil))
	// 分个单独获取
	g:=new(generator.GeneratorData)
	fmt.Println(g.GeneratorPhone())
	fmt.Println(g.GeneratorName())
	fmt.Println(g.GeneratorIDCart())
	fmt.Println(g.GeneratorEmail())
	fmt.Println(g.GeneratorBankID())
	fmt.Println(g.GeneratorAddress())
}

```

# Statement
本项目用于开发环境,涉及商业用途用本人无关
