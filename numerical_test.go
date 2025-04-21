package structure

import (
	"fmt"
	"testing"
	"time"
)

type NumericalDate struct {
	Da *time.Time
	Fl float64
	In int
}

func SCNumericalDate(num int) []NumericalDate {
	var res = make([]NumericalDate, num)
	for i := 0; i < num; i++ {
		in := MathRandom[int](1, 10000)
		fl := MathRandom[float64](1, 10000)
		res[i] = NumericalDate{
			Fl: fl,
			In: in,
		}
	}
	return res
}

func TestNumerical(t *testing.T) {
	var list = SCNumericalDate(10)
	slices := StructField[int](list, "In")
	fmt.Println(Default(SumIndex[int](slices)))
	fmt.Println(Default(AvgIndex[int](slices)))
	fmt.Println(Default(MaxIndex[int](slices)))
	fmt.Println(Default(MinIndex[int](slices)))

	flt := StructField[float64](list, "Fl")
	fmt.Println(Default(SumIndex[float64](flt)))
	fmt.Println(Default(AvgIndex[float64](flt)))
	fmt.Println(Default(MaxIndex[float64](flt)))
	fmt.Println(Default(MinIndex[float64](flt)))

	var slices2 = []*float64{Address[float64](0), Address[float64](10), Address[float64](12), Address[float64](-1)}
	fmt.Println(Default(SumIndex[float64](slices2)))
	fmt.Println(Default(AvgIndex[float64](slices2)))

	//var a = []*int{new(int), new(int), new(int)}
	//a[0] = Address(10)
	//a[0] = Address(10)
	//a[0] = Address(10)
	//
	//fmt.Println(Max(a))

}

//测试 0 值
