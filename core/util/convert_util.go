package util

import (
	"strconv"
)

const DATETIMESTRING = "20060102150405"

func ConvertStringToBoolean(booleanString string, defaultValue bool) bool {
	retBool := defaultValue

	if len(booleanString) > 0 {
		b, err := strconv.ParseBool(booleanString)
		if err == nil {
			retBool = b
		}
	}
	return retBool
}

func ConvertStringToFloat(val string, defaultValue float64) float64 {
	ret := defaultValue

	if len(val) > 0 {
		b, err := strconv.ParseFloat(val, 64)
		if err == nil {
			ret = b
		}
	}
	return ret
}

func ConvertStringToInt(intString string, defaultValue int) int {
	retInt := defaultValue

	if len(intString) > 0 {
		i64, err := strconv.Atoi(intString)
		if err == nil {
			retInt = i64
		}
	}
	return retInt
}

func ConvertStringToInt64(intString string, defaultValue int64) int64 {
	retInt := defaultValue

	if len(intString) > 0 {
		i64, err := strconv.ParseInt(intString, 10, 64)
		if err == nil {
			retInt = i64
		}
	}
	return retInt
}
