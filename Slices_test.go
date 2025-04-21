package structure

import (
	"encoding/json"
	"fmt"
	"log/slog"
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
	res := SC(1000)
	stm := time.Now()
	itr := StructField[string](res, "Address.Phone.Iphone")
	//fmt.Println(itr)
	fmt.Println(len(itr), time.Now().Sub(stm).Seconds())
}
