package tool

import (
	"fmt"
	"github.com/shopspring/decimal"
	"strconv"
)

var oneHundredDecimal decimal.Decimal = decimal.NewFromInt(100)

// Fen2Yuan 分转元
func Fen2Yuan(fen int64) float64 {
	y, _ := decimal.NewFromInt(fen).Div(oneHundredDecimal).Truncate(2).Float64()
	float, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", y), 64)
	return float
}

// Yuan2Fen 元转分
func Yuan2Fen(yuan float64) int64 {

	f, _ := decimal.NewFromFloat(yuan).Mul(oneHundredDecimal).Truncate(0).Float64()
	return int64(f)

}

func Decimal(value float64) float64 {
	value, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", value), 64)
	return value
}
