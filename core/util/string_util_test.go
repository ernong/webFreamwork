package util

import (
	"fmt"
	"testing"
	"time"
)

func TestMD5(t *testing.T) {
	salt := GetRandomString(8)
	pwdCrypt := MD5(salt + "hao123")
	fmt.Println(salt)
	fmt.Println(pwdCrypt)
}

func TestAA(*testing.T) {
	fmt.Print(ConvertStringToFloat("5120.0000", 0))
}

func TestWeek(t *testing.T) {
	thisMon, thisWeek := GetWeekDateOfWeek()
	fmt.Printf("thisMon:%v,thisWeek:%v", thisMon, thisWeek)
}

func TestFormatFloat(t *testing.T) {
	ff := float64(1234.34313214)
	fmt.Print(FormatFloat(ff))
}

func TestTime(t *testing.T) {
	tt := int64(1589877883)
	//reqTs, _ := time.Parse(DATATIMEFORMAT, fmt.Sprintf("%v", tt))
	reqTs := time.Unix(tt, 0)
	ts := reqTs.Format(DATATIMEFORMAT)
	fmt.Printf("req.timestamp:%v, ts:%v", tt, ts)
}
