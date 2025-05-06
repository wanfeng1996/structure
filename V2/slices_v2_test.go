package structure

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"math"
	"math/rand"
	"testing"
	"time"
)

type Address2 struct {
	City  string
	State string
	Phone
}

type Phone struct {
	Iphone string
}

type TestValue struct {
	Name    *string
	Age     *int
	Address Address2
}

func TestGetFieldIndex(t *testing.T) {
	var data = &TestValue{}
	var res []int
	var b bool

	var num = 1000000

	for i := 0; i < num; i++ {
		res, b = GetStructIndex(data, "Address.Phone.Iphone")
	}
	fmt.Println(res, b)

}

type TestValue2 struct {
	Name     *string
	Age      *int
	Address2 // 匿名字段
}

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

func TestSlog(t *testing.T) {
	var data = &TestValue2{
		Name:     nil,
		Age:      nil,
		Address2: Address2{City: "adfa"},
	}
	slog.Info("end", slog.Any("data", func() string {
		marshal, err := json.Marshal(*data)
		if err != nil {
			return ""
		}
		return string(marshal)
	}()))
}

func SC(Num int) []TestValue {
	var res = make([]TestValue, Num)
	var name = "张三"
	for i := 0; i < Num; i++ {
		var age = MathRandom(1, 1000)
		res[i] = TestValue{
			Name: &name,
			Age:  &age,
			Address: Address2{
				Phone: Phone{
					Iphone: "11111111111",
				},
			},
		}
	}

	return res
}

func TestSlices(t *testing.T) {
	res := SC(10000000)
	stm := time.Now()
	itr := StructField(res, "Address.Phone.Iphone")
	//fmt.Println(itr)
	fmt.Println(len(itr), time.Now().Sub(stm).Seconds())
}
