package generator


import "github.com/srlemon/idSDK/metadata"

const (
	ProvinceCityLength = len(metadata.ProvinceCity)
	CardBinsLength =len(metadata.CardBins)
	DomainSuffixLength = len(metadata.DomainSuffix)
	AreaCodeLength = len(metadata.AreaCode)
	CityNameLength = len(metadata.CityName)
	MobilePrefix = len(metadata.MobilePrefix)

)

// GeneratorData 数据
type GeneratorData struct {
	Name string `json:"name"`
	Email string `json:"email"`
	Address string `json:"address"`
	BankID  string `json:"bank_id"`
	IDCard  string `json:"id_card"` // 身份证号
	PhoneNum string `json:"phone_num"` // 手机号码
	// TODO
	// IssueOrg string `json:"issue_org"` // 身份证发证机关
	Birthday string `json:"birthday"` // 出生日期
	ValidPeriod string `json:"valid_period"` // 有效时期
	preCardNo string
}

// NewGeneratorData
func NewGeneratorData()(ret *GeneratorData)  {
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