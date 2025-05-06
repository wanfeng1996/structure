package structure

import (
	"testing"
	"time"
)

type TimeDate struct {
	Da *time.Time
}

func SCTimeDate(num int) []TimeDate {
	var res = make([]TimeDate, num)
	for i := 0; i < num; i++ {
		ttt := MathRandom[int](10, 10000)
		tt := time.Now().Add(time.Duration(ttt) * time.Minute)
		res[i] = TimeDate{
			Da: &tt,
		}
	}
	return res
}

func TestTime(t *testing.T) {
	//var list = SCTimeDate(1000)
	//slices := StructField[time.Time](list, "Da")
	//
	//fmt.Println(transition(slices).Min())
	//fmt.Println(transition(slices).Max())
}
