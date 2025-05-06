package pkg

import (
	"math"
)

func IsNilOrDefault[T comparable](ptr *T) bool {
	if ptr == nil {
		return true
	}
	var zero T
	return *ptr == zero
}

func RoundToTwoDecimals(num float64) float64 {
	return math.Round(num*100) / 100
}

// RoundToDecimal 根据参数保留小数点
func RoundToDecimal(num float64, decimalPlaces int) float64 {
	// 计算 10 的 decimalPlaces 次方
	multiplier := math.Pow(10, float64(decimalPlaces))
	// 四舍五入并返回结果
	return math.Round(num*multiplier) / multiplier
}
