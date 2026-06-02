package main

import (
	"strings"

	"github.com/shopspring/decimal"
)

var ones = []string{
	"", "หนึ่ง", "สอง", "สาม", "สี่",
	"ห้า", "หก", "เจ็ด", "แปด", "เก้า",
}

var positions = []string{
	"", "สิบ", "ร้อย", "พัน", "หมื่น", "แสน", "ล้าน",
}

func convertGroup(n int64) string {
	if n == 0 {
		return ""
	}

	var digits []int64
	tmp := n
	for tmp > 0 {
		digits = append([]int64{tmp % 10}, digits...)
		tmp /= 10
	}

	var parts []string
	for i, d := range digits {
		pos := len(digits) - 1 - i

		if d == 0 {
			continue
		}

		if pos == 1 && d == 2 {
			parts = append(parts, "ยี่สิบ")
			continue
		}

		if pos == 0 && d == 1 && len(digits) > 1 {
			parts = append(parts, "เอ็ด")
			continue
		}

		parts = append(parts, ones[d]+positions[pos])
	}

	return strings.Join(parts, "")
}

func DecimalToThaiText(amount decimal.Decimal) string {
	intPart := amount.Floor()
	fracPart := amount.Sub(intPart).Mul(decimal.NewFromInt(100)).Round(0)

	result := convertGroup(intPart.IntPart()) + "บาท"

	if fracPart.IntPart() == 0 {
		result += "ถ้วน"
	} else {
		result += convertGroup(fracPart.IntPart()) + "สตางค์"
	}

	return result
}
