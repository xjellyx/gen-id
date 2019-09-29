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
func (g *GeneratorData) GeneratorIDCart(isFullAge *bool) (ret *GeneratorData, err error) {
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
	var (
		birthday   string
		code       string
		begin, end time.Time
	)
	if isFullAge == nil {
		isFullAge = new(bool)
		*isFullAge = true
	}
	if t, _err := g.randBirthday(*isFullAge); _err != nil {
		err = _err
		return
	} else {
		g.Birthday = t.UTC().Format("2006-01-02")
		birthday = t.UTC().Format("20060102")
	}

	randomCode := fmt.Sprintf("%0*d", 3, utils.RandInt(0, 999))
	// 合成身份证
	prefix := strconv.Itoa(areaCode) + birthday + randomCode
	if code, err = g.VerifyCode(prefix); err != nil {
		return
	}
	g.IDCard = prefix + code

	// 获取随机有效时间
	if begin, err = g.randDate(); err != nil {
		return
	}
	end = begin.AddDate(10, 0, 0)
	g.ValidPeriod = begin.Format("2006.01.02") + "-" + end.Format("2006.01.02")

	//
	ret = g
	return
}

// randBirthday isFullAge: true 年满18岁
func (g *GeneratorData) randBirthday(isFullAge bool) (ret time.Time, err error) {
	var (
		begin, end time.Time
	)
	if isFullAge {
		if begin, err = time.Parse("2006-01-02 15:04:05", time.Now().AddDate(-70, 0, 0).Format("2006-01-02 15:04:05")); err != nil {
			return
		}
		if end, err = time.Parse("2006-01-02 15:04:05", time.Now().AddDate(-18, 0, 0).Format("2006-01-02 15:04:05")); err != nil {
			return
		}
		ret = time.Unix(utils.RandInt64(begin.UTC().Unix(), end.UTC().Unix()), 0)
	} else {
		if begin, err = time.Parse("2006-01-02 15:04:05", "1970-01-01 00:00:00"); err != nil {
			return
		}
		if end, err = time.Parse("2006-01-02 15:04:05", time.Now().Format("2006-01-02 15:04:05")); err != nil {
			return
		}
		ret = time.Unix(utils.RandInt64(begin.UTC().Unix(), end.UTC().Unix()), 0)
	}

	return
}

// 获取 VerifyCode
func (g *GeneratorData) VerifyCode(cardId string) (ret string, err error) {
	tmp := 0
	for i, v := range metadata.Wi {
		if t, _err := strconv.Atoi(string(cardId[i])); _err == nil {
			tmp += t * v
		} else {
			err = _err
			return
		}
	}
	return metadata.ValCodeArr[tmp%11], nil
}

// randDate 身份证有效期随机时间 有效期最低限制now-10年
func (g *GeneratorData) randDate() (ret time.Time, err error) {
	var (
		begin, end time.Time
	)
	if begin, err = time.Parse("2006-01-02 15:04:05", time.Now().AddDate(-10, 0, 0).Format("2006-01-02 15:04:05")); err != nil {
		return
	}
	if end, err = time.Parse("2006-01-02 15:04:05", time.Now().Format("2006-01-02 15:04:05")); err != nil {
		return
	}
	return time.Unix(utils.RandInt64(begin.Unix(), end.Unix()), 0), err
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
