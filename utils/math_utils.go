package utils

import (
	"github.com/shopspring/decimal"
)

// DivFloat64 float64除法
func DivFloat64(divisor float64, dividend float64) float64 {
	d1 := decimal.NewFromFloat(divisor)
	d2 := decimal.NewFromFloat(dividend)
	res, _ := d1.Div(d2).Float64()
	return res
}

// MultiFloat64 float64乘法
func MultiFloat64(f1 float64, f2 float64) float64 {
	float1 := decimal.NewFromFloat(f1)
	float2 := decimal.NewFromFloat(f2)
	res, _ := float1.Mul(float2).Float64()
	return res
}
