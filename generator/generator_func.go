package generator

import (
	"fmt"
	"github.com/srlemon/gen-id/metadata"
	"github.com/srlemon/gen-id/utils"
	"math"
	"strconv"
	"time"
)

// GeneratorProvinceAdnCityRand 随机获取城市和地址
func (g *GeneratorData) GeneratorProvinceAdnCityRand() (ret string) {
	return metadata.ProvinceCity[utils.RandInt(0, ProvinceCityLength)]
}

// GetAddress 获取地址
func (g *GeneratorData) GeneratorAddress() (ret string) {
	g.Address = g.GeneratorProvinceAdnCityRand() +
		utils.GenRandomLengthChineseChars(2, 3) + "路" +
		strconv.Itoa(utils.RandInt(1, 8000)) + "号" +
		utils.GenRandomLengthChineseChars(2, 3) + "小区" +
		strconv.Itoa(utils.RandInt(1, 20)) + "单元" +
		strconv.Itoa(utils.RandInt(101, 2500)) + "室"
	return g.Address
}

// GetBankID 获取银行卡号
func (g *GeneratorData) GeneratorBankID() (ret string) {
	var (
		// 随机获取卡头
		bank = metadata.CardBins[utils.RandInt(0, CardBinsLength)]
	)
	// 生成 长度 bank.length-1 位卡号
	g.preCardNo = strconv.Itoa(bank.Prefixes[utils.RandInt(0, len(bank.Prefixes))]) + fmt.Sprintf(
		"%0*d", bank.Length-7, utils.RandInt64(0, int64(math.Pow10(bank.Length-7))))
	g.processLUHN()

	return g.BankID
}

// processLUHN 合成卡号
func (g *GeneratorData) processLUHN() {
	checkSum := 0
	tmpCardNo := utils.ReverseString(g.preCardNo)
	for i, v := range tmpCardNo {
		// 数据层确保卡号正确
		tmp, _ := strconv.Atoi(string(v))
		// 由于卡号实际少了一位，所以反转后卡号第一位一定为偶数位
		// 同时 i 正好也是偶数，此时 i 将和卡号奇偶位同步
		if i%2 == 0 {
			// 偶数位 *2 是否为两位数(>9)
			if tmp*2 > 9 {
				// 如果为两位数则 -9
				checkSum += tmp*2 - 9
			} else {
				// 否则直接相加即可
				checkSum += tmp * 2
			}
		} else {
			// 奇数位直接相加
			checkSum += tmp
		}
	}
	if checkSum%10 != 0 {
		g.BankID = g.preCardNo + strconv.Itoa(10-checkSum%10)
	} else {
		// 如果不巧生成的前 卡长度-1 位正好符合 LUHN 算法
		// 那么需要递归重新生成(需要符合 cardBind 中卡号长度)
		g.GeneratorBankID()
	}
}

// GeneratorEmail 生成邮箱
func (g *GeneratorData) GeneratorEmail() (ret string) {
	g.Email = utils.RandStr(8) + "@" + utils.RandStr(5) + metadata.DomainSuffix[utils.RandInt(0, DomainSuffixLength)]
	return g.Email
}

// GeneratorIDCart 生成身份证信息
func (g *GeneratorData) GeneratorIDCart() (ret string, issueOrg, birthday_, validPeriod, addr_ string) {
	// AreaCode
	areaCode := metadata.AreaCode[utils.RandInt(0, AreaCodeLength)]
	// 获取身份证地址
	addr := metadata.IDPrefix[areaCode] + utils.GenRandomLengthChineseChars(2, 3) + "路" +
		strconv.Itoa(utils.RandInt(1, 8000)) + "号" +
		utils.GenRandomLengthChineseChars(2, 3) + "小区" +
		strconv.Itoa(utils.RandInt(1, 20)) + "单元" +
		strconv.Itoa(utils.RandInt(101, 2500)) + "室"
	g.IDCardAddr = addr
	g.IssueOrg = metadata.IDPrefix[areaCode] + "公安局某某分局"
	// 获取随机生日
	t := g.randDate()
	g.Birthday = t.Format("2006-01-02")
	birthday := t.Format("20060102")
	randomCode := fmt.Sprintf("%0*d", 3, utils.RandInt(0, 999))
	// 合成身份证
	prefix := strconv.Itoa(areaCode) + birthday + randomCode
	g.IDCard = prefix + g.VerifyCode(prefix)

	// 获取随机有效时间
	begin := g.randDate()
	end := begin.AddDate(20, 0, 0)
	g.ValidPeriod = begin.Format("2006.01.02") + "-" + end.Format("2006.01.02")
	return g.IDCard, g.IssueOrg, g.Birthday, g.ValidPeriod, g.IDCardAddr
}

// 获取 VerifyCode
func (g *GeneratorData) VerifyCode(cardId string) string {
	tmp := 0
	for i, v := range metadata.Wi {
		t, _ := strconv.Atoi(string(cardId[i]))
		tmp += t * v
	}
	return metadata.ValCodeArr[tmp%11]
}

// TODO 随机时间 1970-2019
func (g *GeneratorData) randDate() time.Time {
	begin, _ := time.Parse("2006-01-02 15:04:05", "1970-01-01 00:00:00")
	end, _ := time.Parse("2006-01-02 15:04:05", "2019-01-01 00:00:00")
	return time.Unix(utils.RandInt64(begin.Unix(), end.Unix()), 0)
}

// GeneratorPhone 生成手机号码
func (g *GeneratorData) GeneratorPhone() (ret string) {
	g.PhoneNum = metadata.MobilePrefix[utils.RandInt(0, MobilePrefix)] + fmt.Sprintf("%0*d", 8, utils.RandInt(0, 100000000))
	return g.PhoneNum
}

// GeneratorName 生成姓名
func (g *GeneratorData) GeneratorName() (ret string) {
	g.Name = metadata.LastName[utils.RandInt(0, len(metadata.LastName))] + metadata.FirstName[utils.RandInt(0, len(metadata.LastName))]
	return g.Name
}
