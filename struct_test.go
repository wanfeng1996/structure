package structure

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"

	"testing"
	"time"
)

func TestIterGetSlice(t *testing.T) {
	//605526 639909
	//10.197.204.19 port=5432 user=gpadmin password=Gzgpadmin123!@# dbname=lte_mr sslmode=disable
	open, err := gorm.Open("postgres", "host=127.0.0.1 port=5432 user=gpadmin password=Gzgpadmin123!@# dbname=lte_mr sslmode=disable")
	if err != nil {
		fmt.Println(err)
		return
	}

	var ws = &WorkOrders{}

	//err = open.Model(&WorkOrder{}).Select("*").Where("id in (?)", []int{658, 633}).Find(ws).Error
	err = open.Model(&WorkOrder{}).Select("*").Limit(10).Find(ws).Error
	if err != nil {
		fmt.Println(err)
		return
	}

	var start = time.Now()

	res := ws.Merge()

	fmt.Println(time.Now().Sub(start).Seconds())
	fmt.Println(res)

}

func TestIterGetMap(t *testing.T) {

}
