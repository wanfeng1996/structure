package structure

import (
	"math"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano()) //随机数时间种子
}

// MathRandom 生成两数之间的随机数 Mustbe min <= max
func MathRandom[T ~int | ~int64 | ~uint | ~uint64 | ~float32 | ~float64](min, max T) T {
	if min == max {
		return min
	}
	//差值
	difference := (max - min) * T(10000)
	t := rand.Int63n(int64(difference) + 1)
	return min + T(unWarp(t, 4))
}

func unWarp(num int64, retain int) float64 {
	return float64(num) / math.Pow10(retain)
}
