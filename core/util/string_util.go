package util

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/big"
	"math/rand"
	"oceanEngineService/bus/entity/errmsg"
	"regexp"
	"strconv"
	"time"
	"unicode"
)

const DATATIMEFORMAT = "2006-01-02 15:04:05"
const DATAFORMAT = "2006-01-02"
const DATAFULLFORMAT = "2006-01-02 00:00:00"
const DATAHOURFORMAT = "2006-01-02 15:00:00"

const CapProfix = float64(10000)
const VProfix = float64(100000000)

func GetRandomString(n int) string {
	const symbols = "0123456789abcdefghjkmnopqrstuvwxyzABCDEFGHJKMNOPQRSTUVWXYZ"
	const symbolsIdxBits = 6
	const symbolsIdxMask = 1<<symbolsIdxBits - 1
	const symbolsIdxMax = 63 / symbolsIdxBits

	prng := rand.New(rand.NewSource(time.Now().UnixNano()))

	b := make([]byte, n)
	for i, cache, remain := n-1, prng.Int63(), symbolsIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = prng.Int63(), symbolsIdxMax
		}
		if idx := int(cache & symbolsIdxMask); idx < len(symbols) {
			b[i] = symbols[idx]
			i--
		}
		cache >>= symbolsIdxBits
		remain--
	}
	return string(b)
}

func MD5(text string) string {
	ctx := md5.New()
	ctx.Write([]byte(text))
	return hex.EncodeToString(ctx.Sum(nil))
}

func IsMailFormat(mail string) bool {
	var mailRe = regexp.MustCompile(`\A[\w+\-.]+@[a-z\d\-]+(\.[a-z]+)*\.[a-z]+\z`)
	return mailRe.MatchString(mail)
}

func NowStr() string {
	tm := time.Unix(time.Now().Unix(), 0)
	return tm.Format("2006-01-02 15:04:05")
}

func GenTimeStr(ts int64) string {
	return time.Unix(ts, 0).Format("2006-01-02")
}

func IsEmptyStr(needCheck string) bool {
	if needCheck != "" && len(needCheck) > 0 {
		return false
	}
	return true
}

//JSONObjectToString 将jason对象转换为string
func JSONObjectToString(v interface{}) (string, error) {
	//检查参数是否有效
	if nil == v {
		return "", fmt.Errorf("no object input")
	}

	var strJSONString string
	buffer, err := json.Marshal(v)
	if err == nil {
		strJSONString = string(buffer)
	}
	return strJSONString, err
}

//String2Float 将string 转换为big.float
func String2Float(val string) (*big.Float, error) {
	ret, ok := new(big.Float).SetString(val)
	if nil == ret || !ok {
		return nil, errmsg.ErrInvalidStringAmount
	}
	return ret, nil
}

func IsHan(str string) bool {
	for _, r := range str {
		if unicode.Is(unicode.Scripts["Han"], r) || (regexp.MustCompile("[\u3002\uff1b\uff0c\uff1a\u201c\u201d\uff08\uff09\u3001\uff1f\u300a\u300b]").MatchString(string(r))) {
			return true
		}
	}
	return false
}

/**
获取本周的日期
*/
func GetWeekDateOfWeek() (weekMonday, weekSunday string) {
	now := time.Now()

	offset := int(time.Monday - now.Weekday())
	if offset > 0 {
		offset = -6
	}

	weekStartDate := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local).AddDate(0, 0, offset)

	weekMonday = fmt.Sprintf("%v 00:00:00", weekStartDate.Format(DATAFORMAT))
	weekSunday = fmt.Sprintf("%v 23:59:59", weekStartDate.Add(6*24*time.Hour).Format(DATAFORMAT))
	return
}

func FormatFloat(val float64) float64 {
	return ConvertStringToFloat(fmt.Sprintf("%.4f", val), 0)
}

func String2Uint64(intStr string) uint64 {
	int64Num := uint64(0)
	if intStr != "" {
		intNum, _ := strconv.Atoi(intStr)
		int64Num = uint64(intNum)
	}
	return int64Num
}
